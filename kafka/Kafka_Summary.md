# Kafka Staff Engineer 文件總覽

這份文件涵蓋 Kafka 架構核心模組、可靠性策略、設定項目與常見設計題，適用於面試或系統架構設計前的知識查核。

---

## 📄 文件目錄與說明

### 01_Kafka_Overview.md  
Kafka 架構總覽、基本元件（Broker、Topic、Partition、ZooKeeper）介紹，Kafka 適合處理的場景與常見優缺點。

### 02_Topic_and_Partition.md  
Topic 與 Partition 的設計邏輯、如何分區、partition 數影響效能與有序性、partition reassignment。

### 03_Producer_Config.md  
Producer 設定參數解析（acks、batch.size、linger.ms、compression）、record 傳送流程、序列化策略。

### 04_Consumer_Group.md  
Consumer group 原理、offset commit、rebalance 策略（Range vs. Sticky）、consumer group 對 HA 的設計關鍵。

### 05_Kafka_idempotence.md  
Producer 的冪等性（idempotence）設計、幾種 message duplication 的來源、如何避免重複投遞。

### 06_Kafka_Reliability.md  
Kafka 的可靠性保證（at-most-once、at-least-once、exactly-once）與對應參數設定與處理策略。

### 07_Replication_and_Durability.md  
副本同步架構（leader/follower）、ISR 機制、min.insync.replicas 設定與持久化的底層保證。

### 08_Failover_and_Recovery.md  
Broker crash、partition leader failover、uncommitted message 處理、Kafka 自動修復與容錯行為。

### 09_Kafka_Broker.md  
Broker 的角色、如何處理 client 連線、訊息 buffer、磁碟儲存原理、index 與 segment 設計。

### 10_Kafka_Stream_and_Connect.md  
Kafka Streams（流處理）與 Kafka Connect（資料整合）的使用場景、stateful operator、sink/source 設定。

### 11_Kafka_vs_Redis_Comparison.md  
Kafka 與 Redis 在訊息處理上的差異、queue 模型比較、可重播性、可靠性、擴充性選擇差異。

---

建議學習順序：  
Kafka_Overview → Producer/Consumer → Topic/Partition → Reliability → Replication → Streams → 比較