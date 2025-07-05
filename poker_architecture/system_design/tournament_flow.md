# 🏆 錦標賽流程說明（MTT）

本文件描述多桌錦標賽（MTT, Multi-Table Tournament）從玩家報名、賽事啟動、桌子配置、遊戲進行、桌子合併與最終結算的完整流程。

---

## 🎯 流程特性

- 📝 賽事尚未開始時，玩家報名資料僅寫入 MongoDB
- 🕒 比賽開始時間一到，由 room_server 自動建立 table_server 並將玩家分配入桌
- 🧾 table_server 會主動推送 `enterRoomRSP` 指令讓玩家進桌
- 🔁 當桌子人數過少，room_server 自動觸發合桌遷移（Table Migration）
- 💥 玩家輸光籌碼後，table_server 會通知 GameRecordService 更新遊戲狀態
- 📋 玩家可透過 `tableInfoREQ` 查詢目前桌子狀況
- 🏁 最終由 room_server 統一整理名次與獎勵分配

---

## 🔁 MTT 錦標賽完整流程圖

```mermaid
sequenceDiagram
    participant Client as 玩家 Client
    participant Gateway as Game Gateway
    participant GameRouter as GameRouter (快取路由)
    participant Room as room_server (錦標賽管理)
    participant Table as table_server (遊戲進行)
    participant Record as GameRecordService
    participant Mongo as MongoDB

    %% 1. 玩家報名，但比賽尚未開始
    Client->>Gateway: iJoinREQ(tourney_id)
    Gateway->>GameRouter: Forward iJoinREQ
    GameRouter->>Room: iJoinTournamentREQ(uid)
    Room->>Mongo: Insert tournament_enroll(uid, tourney_id)
    Room-->>Client: joinRSP(success)

    %% 2. 比賽開始，由 room_server 配桌
    Note over Room: 比賽時間觸發配桌
    Room->>Mongo: 查詢 tournament_enroll 名單
    loop 每 9 人配一桌
        Room->>Table: iCreateTable(table_id, seat_players)
        Table->>Gateway: enterRoomRSP(table_id, seat_no, uid)
        Room->>Mongo: 寫入 playing_room_status
    end

    %% 3. 玩家操作 Bet / Fold 等
    Client->>Gateway: BetREQ / FoldREQ
    Gateway->>GameRouter: Forward iBetREQ
    GameRouter->>Table: iBetREQ

    %% 4. 玩家輸光籌碼，被淘汰
    Note over Table: 玩家 chip=0 → 淘汰出局
    Table->>Record: UpdatePlayerState(uid, status="out")
    Record->>Mongo: Update playing_room_status / tournament_state

    %% 5. 單桌結果回傳給 room_server
    Table-->>Room: TableResult(table_id, chip_diff)
    Room->>Mongo: Update tournament_state

    %% 6. 桌子人數過少，自動合桌遷移
    alt 桌子人數 < 閾值
        Room->>Table: iMigratePlayers(from → to)
        Table-->>Room: PlayerMigrateACK
        Room->>Mongo: Update playing_room_status
    end

    %% 7. 玩家查詢 tableInfo
    Client->>Gateway: tableInfoREQ(table_id)
    Gateway->>GameRouter: Forward tableInfoREQ
    GameRouter->>Room: iTableInfoREQ(table_id)
    Room->>Mongo: Read playing_room_status
    Room-->>Client: tableInfoRSP(players, chip, blinds, level...)

    %% 8. 比賽結束與結算
    alt 僅剩一桌 or 排名完成
        Table-->>Room: FinalResult(rankings)
        Room->>Mongo: Insert tournament_result
        Room-->>Client: TournamentEndRSP(rank, prize)
    end
```

## 🔁 桌子合併細節流程（TableA → TableB）
```mermaid
sequenceDiagram
    participant Room as room_server
    participant TableA as table_server A
    participant TableB as table_server B
    participant Redis as Redis (可選)
    participant Mongo as MongoDB

    %% 1. 偵測到 TableA 人數低於閾值（如 < 3）
    Note over Room: 偵測 TableA 人數不足

    %% 2. Room 發送指令，準備遷移至 TableB
    Room->>TableA: iPrepareMigrate(table_id=B)

    %% 3. TableA 回傳待遷移玩家列表
    TableA-->>Room: PlayerList(uid_list)

    %% 4. Room 發送移動指令給 TableB
    Room->>TableB: iMigratePlayers(uid_list)

    %% 5. TableB 執行 addPlayer，準備座位與狀態同步
    TableB-->>Room: PlayerMigrateACK

    %% 6. Room 通知 TableA 清理該桌（或標記為待清除）
    Room->>TableA: iReleaseTable

    %% 7. Room 更新 Mongo 的桌子狀態資料
    Room->>Mongo: Update playing_room_status
```