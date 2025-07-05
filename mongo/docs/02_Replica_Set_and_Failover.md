
# MongoDB Replica Set：Failover 機制解析

本章節說明 MongoDB 在 Replica Set 中的自動故障轉移（failover）機制，幫助你理解其高可用性設計原理，並掌握選舉流程與常見參數的運作。

---

## 🔁 Replica Set 架構概念

- 一組 `mongod` 節點組成 Replica Set，通常有 **1 Primary + N Secondary**
- 所有寫入僅發生在 Primary；讀取可配置從 Secondary 讀（視 Read Preference）
- Primary 若故障，會由 Secondary 自動選舉新的 Primary

```
[Primary] ←→ [Secondary 1]
     ↑        ↓
[Secondary 2]
```

---

## ⚡ Failover 發生時的流程

1. Primary 節點異常，停止心跳響應
2. 其餘節點發現 Primary 失聯（約 10 秒）
3. 剩餘節點自動觸發 **選舉（election）**
4. 若有過半數投票成功，新的 Primary 成立
5. 所有節點更新其角色與資料同步方向

---

## 🧠 選舉流程與規則

| 規則                         | 說明 |
|------------------------------|------|
| 過半數原則                   | 選舉成功需取得超過半數節點同意 |
| 投票數最多為 7               | 即便有 50 節點，最多只會有 7 票 |
| `priority` 參數控制選舉權重 | 可設定特定節點優先成為 Primary |
| Arbiter 只參與投票           | 不儲存資料，但幫助達成過半 |

---

## ⚙️ 常見設定參數

| 參數名稱        | 說明 |
|------------------|------|
| `priority`       | 設定節點優先成為 Primary 的權重（0 表示永不成為 Primary）|
| `votes`          | 每個節點的投票權重（最多 1）|
| `hidden`         | 隱藏節點，不被用於查詢或選舉 |
| `buildIndexes`   | 是否參與建立索引 |
| `arbiterOnly`    | 是否為仲裁節點（無資料，只投票）|

---

## 🧪 實務建議

- 最佳節點數為奇數（例如：3、5、7）→ 易於過半數決
- 若需加入備援但不希望參與選舉，可設 `priority: 0`
- Arbiter 只能在少量資料或小型集群中使用，避免成為單點風險
- 使用監控系統（如 MongoDB Ops Manager、Prometheus）追蹤 Primary 狀態

---

## 🧠 Staff Engineer 該理解的重點

- MongoDB 的 HA 設計不需外部 coordinator，可自動完成選舉
- Failover 雖自動，但選舉期間會短暫中斷寫入（應設計 retry 機制）
- 與 CAP 理論相關，Mongo Replica Set 為 AP 系統，分區時仍允許讀寫但需謹慎設計

---

[← 回到總覽](../Mongo_Summary.md)
