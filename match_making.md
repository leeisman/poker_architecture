# Matchmaking 自動配桌流程

當玩家進入房間（EnterRoomREQ）後，RoomServer 將根據當前桌子使用狀況、自訂條件（如 VIP、盲注等）進行自動配桌。此流程在 Ring Game 中尤為關鍵，可確保桌子資源最佳化使用並減少等待時間。

---

## 🎯 設計目標

- 動態分配玩家至已有桌子（優先坐滿）
- 若無可用座位，自動建立新桌子
- 保證玩家體驗順暢、支援大量併發進入

---

## 🌐 流程圖（Matchmaking）

```mermaid
sequenceDiagram
    participant Client as 玩家 Client
    participant Gateway as Game Gateway
    participant Router as GameRouter（指令路由）
    participant Room as RoomServer（房間管理）
    participant Table as TableServer（遊戲邏輯）
    participant DB as MongoDB（playing_room_status）

    Client->>Gateway: EnterRoomREQ
    Gateway->>Router: EnterRoomREQ
    Router->>Room: iEnterRoomREQ

    %% 配桌邏輯
    Room->>DB: 查詢 playing_room_status
    alt 有可用座位
        Room->>Table: iSitDownREQ（加入現有桌）
    else 沒有可用桌
        Room->>Table: iCreateTableREQ
        Table->>DB: 初始化 table 狀態
        Room->>Table: iSitDownREQ（加入新桌）
    end

    %% 最終回應
    Table-->>Gateway: EnterRoomRSP
    Gateway-->>Client: EnterRoomRSP
```