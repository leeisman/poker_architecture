# 11. Retry Strategy and Error Handling

本章節探討 MongoDB 在錯誤處理與重試策略上的應用。針對網路錯誤、主從切換、交易衝突等場景，正確的重試機制是穩定服務不可或缺的一環。

---

## 🔁 一、MongoDB 驅動支援的自動重試機制

### ✅ 支援的操作
- `insertOne`, `updateOne`, `deleteOne`（單筆）
- `findOneAndUpdate`, `findOneAndDelete`
- 多數 read 操作支援 retryable read

### ⚙️ 啟用條件
- 使用 MongoDB 4.2+，且連線為 replica set 或 sharded cluster
- 驅動層需設定 retryWrites=true（預設為開）

```go
clientOptions := options.Client().ApplyURI("mongodb://host/?retryWrites=true")
```

---

## ⚠️ 二、常見錯誤類型與處理方式

| 錯誤代碼 / 類型     | 說明                          | 建議處理方式           |
|----------------------|-------------------------------|------------------------|
| `NetworkError`       | 連線中斷、超時等               | 可進行重試             |
| `NotPrimary`         | 原主節點已轉為備節點           | 重新選主 + 重試        |
| `WriteConflict`      | 多筆操作發生衝突               | 延遲後 retry           |
| `DuplicateKey`       | 違反唯一索引約束               | 業務判斷是否忽略       |
| `TransactionAborted` | 複雜交易執行失敗               | 適度重試               |
| `StaleShardVersion`  | shard metadata 過期            | mongos 會自動更新重試  |

---

## 🔃 三、實作 retry 時的注意事項

- **需確保操作具備重試安全性（idempotent）**
- 設定最大 retry 次數與 backoff 機制（exponential backoff）
- 寫入操作儘可能採用 `writeConcern` = majority 避免副本不一致
- 若搭配交易，需搭配 `txnNumber` 管理

---

## 🛡 四、與交易搭配時的 retry 模式（Go 範例）

```go
for retries := 0; retries < 3; retries++ {
    sess, err := client.StartSession()
    if err != nil { ... }

    err = mongo.WithSession(ctx, sess, func(sc mongo.SessionContext) error {
        err := sess.StartTransaction()
        if (err != nil) { return err }

        err = collection.UpdateOne(sc, ...)
        if err != nil {
            sess.AbortTransaction(sc)
            return err
        }

        return sess.CommitTransaction(sc)
    })

    if err == nil {
        break
    }

    time.Sleep(time.Second * time.Duration(retries+1))
}
```

---

## 🧠 Staff Engineer 該理解的點

- MongoDB 的 retry 機制由 driver 管理，需搭配正確連線設定與 error handling
- 僅有「單筆操作」才支援 automatic retry
- 複雜操作建議加上應用層級重試與幂等邏輯設計
- 須了解常見錯誤代碼與其對應重試策略

---

[← 回到總覽](../Mongo_Summary.md)