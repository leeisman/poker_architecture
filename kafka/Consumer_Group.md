# 👥 Kafka Consumer Group 詳解

Kafka 中的 **Consumer Group** 是實現**水平擴展、容錯處理**與**可追蹤性**的核心機制之一。本文件將說明其運作方式、設計優勢與實務應用場景。

---

## 🧠 基本概念

- **Consumer Group** 是由多個 Consumer 組成的群組，用來**共同處理某個 Topic 的訊息**。
- **Partition 是不可共享的**，在同一個 Consumer Group 裡，一個 Partition 同一時間只能被一個 Consumer 消費。
- Kafka 保證 **每筆訊息只被 Group 中的一個 Consumer 消費一次**。

---

## 🔁 Consumer Group 如何分配 Partition

```mermaid
flowchart TD
    subgraph Topic: user_events
        P1[Partition 0]
        P2[Partition 1]
        P3[Partition 2]
    end

    subgraph Consumer Group: user_event_processors
        C1[Consumer 1]
        C2[Consumer 2]
    end

    P1 --> C1
    P2 --> C2
    P3 --> C1
```

---

```mermaid
sequenceDiagram
    participant Consumer1
    participant Consumer2
    participant Broker1 as Broker 1 (Group Coordinator)

    Note over Consumer1,Consumer2: 加入同一個 Consumer Group "order-group"

    Consumer1->>Broker1: JoinGroupRequest
    Consumer2->>Broker1: JoinGroupRequest

    Broker1-->>Consumer1: Partition 0, 1 分配
    Broker1-->>Consumer2: Partition 2 分配

    Consumer1->>Broker1: Heartbeat (定期回報狀態)
    Consumer2->>Broker1: Heartbeat
```

--- 

## 同一 Group 同時兩個 Consumer consume  Partition 流程


```mermaid
sequenceDiagram
    participant ZK as Zookeeper
    participant Broker as Kafka Broker
    participant C1 as Consumer 1（先啟動）
    participant C2 as Consumer 2（後啟動）
    participant P0 as Partition 0
    participant P1 as Partition 1
    participant P2 as Partition 2
    participant P3 as Partition 3

    %% Consumer 1 啟動流程
    C1->>Broker: Join Group (group.id="my-group")
    Broker-->>ZK: 註冊 Consumer 1 in group
    ZK-->>Broker: OK
    Broker-->>C1: 分配 Partition 0,1,2,3

    %% Consumer 2 加入，觸發 rebalance
    C2->>Broker: Join Group (group.id="my-group")
    Broker-->>ZK: 更新 Group Metadata
    Note over Broker,ZK: Broker 發現 Group 成員變化 → 觸發 Rebalance

    %% Rebalance 分配
    Broker-->>C1: 分配 Partition 0,1
    Broker-->>C2: 分配 Partition 2,3

    %% 消費流程重新啟動
    C1->>P0: Fetch messages
    C1->>P1: Fetch messages
    C2->>P2: Fetch messages
    C2->>P3: Fetch messages
```