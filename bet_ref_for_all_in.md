# 雙方 All-In（BetREQ）處理流程說明

當兩名玩家（A 與 B）在同一局遊戲中進行 All-In，系統將啟動完整的下注、發牌、勝負判定、紀錄寫入、廣播等流程。本文件描述此流程中每個模組的責任範疇與異常處理。

---

## 🎯 架構設計重點

- 所有狀態由 `table_server` 控制並具有內部狀態機
- 路由與服務發現交由 `game_router` 管理
- 資料寫入經由 `game_record_server` 統一處理，避免 table_server 操作資料庫造成 I/O 過重

---

## 🧩 All-In 處理流程（BRC 模型）

```mermaid
sequenceDiagram
    participant Client1 as 玩家 A
    participant Client2 as 玩家 B
    participant Gateway as Game Gateway
    participant Router as game_router<br/>(查找 table instance)
    participant Table as table_server
    participant Record as game_record_server
    participant Mongo as MongoDB

    %% 玩家 A、B All-In 發起請求
    Client1->>Gateway: BetREQ (All-In)
    Client2->>Gateway: BetREQ (All-In)

    %% Gateway 轉交至 Router（指令路由）
    Gateway->>Router: Forward iBetREQ (A)
    Gateway->>Router: Forward iBetREQ (B)

    %% Router 查表位，找到 table_server 實體，直接轉發
    Router->>Table: iBetREQ (A)
    Router->>Table: iBetREQ (B)

    %% table_server 成功處理，直接回應 Gateway
    Table-->>Gateway: BetRSP (A)
    Table-->>Gateway: BetRSP (B)
    Gateway-->>Client1: BetRSP
    Gateway-->>Client2: BetRSP

    %% B: Broadcast 發牌流程（翻牌 Turn River）
    Table->>Gateway: Broadcast DealCards
    Gateway->>Client1: DealCards
    Gateway->>Client2: DealCards

    %% R: Resolve 勝負 + Side Pot 分配
    Table->>Table: Evaluate Hands & SidePot

    %% C: Commit 結果 → 寫入紀錄服務
    Table->>Record: HandEndREQ
    Record->>Mongo: Insert game_record, Update gameset_record

    %% 回傳結果
    Table->>Gateway: Broadcast ShowdownResult
    Gateway->>Client1: ShowdownResult
    Gateway->>Client2: ShowdownResult
```