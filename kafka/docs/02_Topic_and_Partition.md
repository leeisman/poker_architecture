# 🧩 Kafka Topic & Partition 概念說明

## 📦 Topic 是什麼？

在 Kafka 中，**Topic 是訊息的邏輯分類單位**，所有的訊息都被發送到某個特定的 Topic，例如：

- `user_signup`
- `order_events`
- `clickstream`

---

## 🧱 Partition 是什麼？

每個 Topic 被分割為多個 **Partition**，Partition 是 Kafka 的水平擴展與並行處理基礎：

- 每個 Partition 是一條 **有序的 append-only log**
- 訊息在 Partition 中會依照寫入順序被編號（offset）
- 不同 Partition 可由不同 broker/consumer 處理，實現高併發

---

## 🧠 為什麼要分 Partition？

| 目的         | 好處                                                 |
|--------------|------------------------------------------------------|
| 高可用       | 不同 Partition 可分佈到不同 Kafka Broker           |
| 水平擴展     | 較多 Consumer 可平行處理訊息，提高吞吐量           |
| 提高容錯能力 | Partition 可設 Replica，某台 Broker 掛掉也能恢復   |

---

## 🎯 舉例說明

假設一個 Topic 有 3 個 Partition，而一個 Consumer Group 有 2 個 Consumer：

```mermaid
flowchart TD
    subgraph Topic: order_events
        P1[Partition 0]
        P2[Partition 1]
        P3[Partition 2]
    end

    subgraph Consumer Group: order_processors
        C1[Consumer 1]
        C2[Consumer 2]
    end

    P1 --> C1
    P2 --> C2
    P3 --> C1
```
---
## Broker & Partition

```mermaid
flowchart TD
    subgraph Kafka Cluster
        B1[Broker 1]
        B2[Broker 2]
        B3[Broker 3]
        T1_P0[(TopicA - Partition 0)]
        T1_P1[(TopicA - Partition 1)]
    end

    B1 -->|Leader_P0| T1_P0
    B2 -->|Follower| T1_P0
    B3 -->|Follower| T1_P0

    B2 -->|Leader_P1| T1_P1
    B1 -->|Follower| T1_P1
    B3 -->|Follower| T1_P1
```
---
```mermaid
sequenceDiagram
  participant Producer0
  participant Broker1
  participant Broker2
  participant Broker3
  participant Controller (in Broker2)

  Producer0->>Broker1: MetadataRequest(TopicA)
  Broker1-->>Producer0: P0 → B2, P1 → B1, P2 → B3, P3 → B1

  Note right of Controller: Controller 維護 leader map

  Producer0->>Broker2: Send record to P0
  Producer0->>Broker1: Send record to P1
  Producer0->>Broker3: Send record to P2
  Producer0->>Broker1: Send record to P3
```