# Zookeeper 介紹與 Golang 中的服務發現實作

## 🔍 Zookeeper 是什麼？
Zookeeper 是一個由 Apache 提供的 **分散式協調服務**，主要用來：

- 維護分散式系統中的配置信息
- 命名服務
- 分布式鎖
- 領導者選舉
- **服務發現**

Zookeeper 具有以下特性：
- 高可用性（透過 ZooKeeper Ensemble）
- 強一致性（ZAB 協議）
- 原子性、順序性操作
- 支援 Watcher 機制（可觸發事件）

---

## 📦 Zookeeper 資料結構：ZNode
Zookeeper 中的資料存放在類似檔案系統的結構裡：
/
├── services
│   └── game
│       ├── 192.168.0.1:8080   <- Ephemeral node
│       └── 192.168.0.2:8080

---
ZNode 分為：
- **Persistent Node**：永久存在，除非手動刪除
- **Ephemeral Node**：臨時節點，client 斷線後自動刪除

---