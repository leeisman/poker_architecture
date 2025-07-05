# 分散式系統設計文件總覽（Staff Engineer 準備）

本專案資料夾收錄各領域在面試或系統設計中常用的核心技術主題，適合用於整理學習、架構準備、或內部知識分享。

---

## 📁 資料夾說明

### ✅ golang/
- Goroutine、Channel、Map、Sync、Memory、Fuzzing、性能分析、進階實戰

### ✅ kafka/
- Kafka 核心架構、Topic/Partition、Producer/Consumer 設定、可靠性、Broker、Stream、Redis 比較

### ✅ mongo/
- Replica Set、Sharding、Aggregation、Schema Design、交易一致性、Read/Write Concern、性能與錯誤處理

### ✅ mysql/
- 索引設計、Query Plan 分析、主從複製、鎖機制、JOIN 策略、交易隔離級別、分頁與分表技巧

### ✅ network/
- TCP/IP、HTTP/2、gRPC、負載平衡、Connection Pool、KeepAlive、網路延遲與併發連線瓶頸

### ✅ poker_architecture/
- 德州撲克遊戲的分散式系統設計，涵蓋配桌、賽事、ALL-IN、超時恢復、Zookeeper、負載平衡等模組

### ✅ redis/
- Redis 基礎、Cluster 架構、寫入策略、Stream、Pub/Sub、Cache Aside、過期與淘汰策略

### ✅ zookeeper/
- ZooKeeper 原理、Ephemeral node、Watch、CAP、Leader 選舉、分散式鎖、實作挑戰經驗

---

## 📌 使用建議

- 適用角色：後端工程師、架構師、Staff Engineer 應徵者
- 搭配實戰練習與 Mock Interview 可快速補強模組思維
- 可結合各語言實作範例 (Golang 為主)