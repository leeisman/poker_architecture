# EnterRoomREQ 進房請求處理流程

本文件描述玩家進入德州撲克房間時的整體流程，涵蓋指令傳遞、房間路由、座位分配與回應邏輯。透過 GameRouter 解耦 Gateway 與 RoomServer，並配合服務發現與資料持久化達成高可用與可觀察性。

---

## 🎯 架構設計目標

- 🧭 GameRouter 負責進房指令統一入口與服務定位
- 🏠 RoomServer 進行座位邏輯、與 Table 協同工作
- 🧠 Mongo 為房間狀態的持久化來源
- 🎲 Table 服務專注於遊戲狀態與玩家管理
- 🎯 玩家最終由 Table 服務回應結果

---

## 🧩 請求流程圖

```mermaid
sequenceDiagram
    participant Client as 玩家 Client
    participant Gateway as Game Gateway<br/>(WebSocket)
    participant GameRouter as GameRouter<br/>(服務發現 + 轉發)
    participant Mongo as MongoDB
    participant ZK as Zookeeper
    participant Room as RoomServer<br/>(房間管理)
    participant Table as TableService<br/>(發牌 + 結算 + 狀態)

    %% 玩家連線並送出請求
    Client->>Gateway: WebSocket連線建立
    Client->>Gateway: EnterRoomREQ

    %% Gateway 轉發請求給 GameRouter
    Gateway->>GameRouter: EnterRoomREQ 轉發

    %% GameRouter 查詢 Mongo 確認房間資訊
    GameRouter->>Mongo: 查詢 room_id 對應的實例與設定

    %% GameRouter 透過 Zookeeper 找到 RoomServer 實例
    GameRouter->>ZK: 查詢 room_xxx 服務位址
    GameRouter->>Room: 將 iEnterRoomREQ 轉交至 RoomServer

    %% RoomServer 處理座位分配並建立/指派 Table
    Room->>Table: iSitDownREQ（包含 uid、seat_no、room_id）

    %% Table 回傳入桌資訊（如 table_id, seat_no）
    Table-->>Room: 回傳座位結果
    Room->>Mongo: 更新 playing_room_status 文件（更新內部記憶體資料至持久層）

    %% ✅ TableService 最終直接回傳結果給 Gateway
    Table-->>Gateway: EnterRoomRSP
    Gateway-->>Client: EnterRoomRSP
```