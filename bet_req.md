# BetREQ 請求處理流程

本文件描述玩家下注（BetREQ）時，如何透過 Gateway → GameRouter → TableServer 的分層架構進行路由、快取查找與服務發現。目的是在保證高可用下，快速且正確將指令導入對應的遊戲執行實例（TableServer）。

---

## 🎯 架構設計目標

- ❄️ Gateway 輕量處理：不保管任何狀態，只負責收發封包
- 🚦 GameRouter 多層緩存：支援 Cache / MongoDB fallback / Zookeeper 查找
- 🎲 TableServer 為唯一遊戲邏輯執行單位
- 🔁 支援 tableServer crash 自我修復與 client 端 retry

---

## 🧩 請求流程圖

```mermaid
sequenceDiagram
    participant Client as 玩家 Client
    participant Gateway as WebSocket Gateway
    participant Router as GameRouter<br/>(查詢 tableInstance 快取 + 路由)
    participant Cache as 記憶體快取 (table routing)
    participant Mongo as MongoDB（playing_room_status）
    participant ZK as Zookeeper（服務發現）
    participant Table as TableServer（遊戲邏輯）

    %% 玩家送出請求
    Client->>Gateway: BetREQ(room_id, table_id)

    %% Gateway 將請求轉給 GameRouter 處理
    Gateway->>Router: Forward BetREQ(room_id, table_id)

    %% Router 查詢 Cache
    Router->>Cache: 查詢 tableInstance

    alt Cache Miss
        Router->>Mongo: 查找 playing_room_status（room_id + table_id → tableInstance）

        alt 找到 tableInstance
            Router->>ZK: 查找 TableServer 服務位址

            alt 找到服務
                Router->>Table: BetREQ
                Table-->>Router: BetRSP
                Router-->>Gateway: BetRSP
                Gateway-->>Client: BetRSP
            else Table 未註冊
                Note right of Router: 不回應（由 Client 處理重試）
            end

        else Mongo 沒有資料
            Note right of Router: 不回應（由 Client 處理重試）
        end

    else Cache Hit
        Router->>ZK: 查找 TableServer 位址
        Router->>Table: BetREQ
        Table-->>Router: BetRSP
        Router-->>Gateway: BetRSP
        Gateway-->>Client: BetRSP
    end
```