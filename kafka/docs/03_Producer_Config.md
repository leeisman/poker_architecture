# 🛠 Kafka Producer Config 設定指南

本文件說明 Kafka Producer 常用的設定選項與其效能、可靠性上的意涵，幫助你根據業務需求調整參數以達到最佳效能與穩定性。

---

## ⚙️ 核心設定參數說明

| 參數名稱            | 預設值     | 說明                                                                 |
|---------------------|------------|----------------------------------------------------------------------|
| `acks`              | `1`        | **回應等級**：`0`, `1`, `all`，決定需等待幾個 Broker 回應才視為成功 |
| `retries`           | `0`        | **重試次數**：訊息傳送失敗時的最大重試次數                         |
| `retry.backoff.ms`  | `100`      | **重試間隔時間**（毫秒）                                             |
| `batch.size`        | `16384`    | 一批訊息的最大總 byte 數（不是筆數）                               |
| `linger.ms`         | `0`        | **延遲時間**（毫秒），Producer 等待更多訊息來湊 batch               |
| `compression.type`  | `none`     | 壓縮方式：`none`, `gzip`, `snappy`, `lz4`, `zstd`                   |
| `buffer.memory`     | `33554432` | Producer buffer 可用記憶體（byte）                                  |
| `max.in.flight.requests.per.connection` | `5` | 同時尚未回應的 request 數量                                          |
| `enable.idempotence`| `false`    | 啟用後避免重複送出（Exactly-once）                                  |
| `key.serializer`    | 必填       | Key 的序列化器，如：`StringSerializer`                              |
| `value.serializer`  | 必填       | Value 的序列化器，如：`StringSerializer`                            |

---

## ✅ 常見設定組合建議

# 🚀 Kafka Producer Configuration 說明（高吞吐量配置）

本文件說明 Kafka Producer 的重要參數設定，並特別介紹在 **高吞吐量場景** 中的推薦配置與設計原理。

---

## ⚙️ 高吞吐量建議設定

| 參數               | 推薦值     | 說明                                                                 |
|--------------------|------------|----------------------------------------------------------------------|
| `acks`             | `1`        | 等 Leader 收到訊息就回應，提升速度，但可能損失資料（相對於 `all`） |
| `batch.size`       | `32768`    | 增加批次大小（bytes），一次傳更多訊息，減少 I/O 次數                 |
| `linger.ms`        | `10`       | 最多等待 10ms，即使 batch 沒滿也送出，加強 batch 機會               |
| `compression.type` | `lz4`      | 使用壓縮減少網路與磁碟負擔，lz4 是壓縮效率與效能的平衡選擇          |

---

## 🎯 適用場景

- 訊息量大（如：event logging、使用者行為追蹤）
- 可接受偶發丟訊息
- 延遲敏感度不高

---

## 🧪 `acks` 行為示意圖

```mermaid
sequenceDiagram
    participant Producer
    participant BrokerL as Leader Broker
    participant BrokerR1 as Replica 1
    participant BrokerR2 as Replica 2

    Note over Producer: 發送訊息 msg #1 給某個 partition

    Producer->>BrokerL: produce request (msg #1)

    alt acks = 0
        Note over Producer: 不等待回應，直接完成
    else acks = 1
        BrokerL-->>Producer: 回傳 ACK（Leader 寫入成功）
    else acks = all
        BrokerL->>BrokerR1: 同步訊息
        BrokerL->>BrokerR2: 同步訊息
        BrokerR1-->>BrokerL: 確認同步完成
        BrokerR2-->>BrokerL: 確認同步完成
        BrokerL-->>Producer: 所有 ISR 完成 → 回 ACK
    end
```
---
## 平行處理
```mermaid
sequenceDiagram
    participant GoApp as Golang App
    participant Producer as Kafka Producer
    participant ZK as Zookeeper
    participant Controller as Kafka Controller
    participant Broker0 as Broker for P0
    participant Broker1 as Broker for P1
    participant Broker2 as Broker for P2
    participant Broker3 as Broker for P3
    participant Broker4 as Broker for P4

    %% Metadata 查詢
    GoApp->>Producer: Send("order_id=xyz")
    Producer->>ZK: Request metadata (TopicA)
    ZK-->>Producer: Partitions info for TopicA
    Producer->>Controller: Request Leader info
    Controller-->>Producer: P0 → Broker0
    Controller-->>Producer: P1 → Broker1
    Controller-->>Producer: P2 → Broker2
    Controller-->>Producer: P3 → Broker3
    Controller-->>Producer: P4 → Broker4

    %% Producer 分流發送訊息
    Producer->>Broker0: Produce(topic=TopicA, partition=0, msg0)
    Broker0-->>Producer: ACK(msg0 成功)

    Producer->>Broker1: Produce(topic=TopicA, partition=1, msg1)
    Broker1-->>Producer: ACK(msg1 成功)

    Producer->>Broker2: Produce(topic=TopicA, partition=2, msg2)
    Broker2-->>Producer: ACK(msg2 成功)

    Producer->>Broker3: Produce(topic=TopicA, partition=3, msg3)
    Broker3-->>Producer: ACK(msg3 成功)

    Producer->>Broker4: Produce(topic=TopicA, partition=4, msg4)
    Broker4-->>Producer: ACK(msg4 成功)
```