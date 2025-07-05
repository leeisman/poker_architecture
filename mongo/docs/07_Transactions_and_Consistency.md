# 07. Transactions and Consistency

MongoDB 自 4.0 起正式支援多文件（multi-document）交易，補足了過去 NoSQL 在強一致性上的短板。作為 Staff Engineer，了解交易模型與一致性保證對系統可靠性設計至關重要。

---

## 🔐 為什麼需要交易（Transaction）？

- 保證多個操作要「全部成功」或「全部失敗」
- 防止資料在部分寫入後造成資料不一致
- 特別適合金融轉帳、庫存扣減、複雜更新

---

## ✅ MongoDB 的交易特性

| 特性             | 說明                                   |
|------------------|----------------------------------------|
| 多文件支援       | 可一次操作多個 collection              |
| 跨 shard 支援    | 自 4.2 起可跨分片交易                   |
| ACID 保證        | 提供原子性（Atomicity）與一致性         |
| Replica Set 必備 | 僅支援在 Replica Set 或 Sharded Cluster 上使用 |

---

## 🧪 使用範例（Golang）

```go
session, err := client.StartSession()
if err != nil { log.Fatal(err) }

defer session.EndSession(context.TODO())

result, err := session.WithTransaction(context.TODO(), func(sc mongo.SessionContext) (interface{}, error) {
    collection1 := client.Database("test").Collection("accountA")
    collection2 := client.Database("test").Collection("accountB")

    if _, err := collection1.UpdateOne(sc, filter1, update1); err != nil {
        return nil, err
    }
    if _, err := collection2.UpdateOne(sc, filter2, update2); err != nil {
        return nil, err
    }
    return nil, nil
})
```

---

## 🔍 一致性模型

MongoDB 的預設一致性為「最終一致性」，但透過以下機制可提升一致性保證：

| 機制             | 描述                                      |
|------------------|-------------------------------------------|
| Write Concern    | 控制寫入幾個節點才算成功（如 `"majority"`）|
| Read Concern     | 控制從哪個節點讀取（如 `"local"`, `"majority"`）|
| Journaling       | 控制是否寫入磁碟 journal（`j: true`）      |

---

## 💡 常見交易錯誤處理

- **TransientTransactionError**：可重試
- **UnknownTransactionCommitResult**：commit 成功與否未知 → 也應重試
- 建議搭配 retry loop 實作交易穩定性

---

## 🧠 Staff Engineer 該理解的點

| 項目                     | 原因                                         |
|--------------------------|----------------------------------------------|
| 跨 collection 的交易風險 | 可能影響效能與資源鎖定                      |
| writeConcern + readConcern 配合 | 寫入與讀取一致性配置需謹慎調整       |
| 避免長交易               | 長時間交易會造成鎖定、blocking，需拆分優化  |

---

## 🚦 Sharded Cluster 補充

在 Sharded Cluster 上使用交易需注意：

- 交易需透過 `mongos` 路由器進行
- 所有涉及的 shard 需參與交易 → 增加 overhead
- 寫入集中在單一 shard 可減少交易成本

---

[← 回到總覽](../Mongo_Summary.md)
