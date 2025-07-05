# 13. Mongo Design Questions

本章節整理常見 MongoDB 設計相關問題，涵蓋 schema 設計、擴充策略、資料一致性、效能調整與實務選型等，協助在面試或實務中具備清晰思考架構。

---

## 🔸 1. 你會如何設計一個聊天系統的資料結構？

### 🎯 考點
- 使用 embedding vs referencing
- 訊息量大時如何分片、分表
- 是否支援多聊天室、歷史訊息查詢

### 💡 建議方向

```json
{
  "_id": ObjectId,
  "chatRoomId": "room-123",
  "messages": [
    { "sender": "user1", "text": "Hi", "timestamp": ISODate },
    { "sender": "user2", "text": "Hello", "timestamp": ISODate }
  ]
}
```

- 若訊息過多可考慮獨立 messages collection（referencing）
- 需查詢近期訊息 → 建 index on chatRoomId + timestamp
- 可搭配 TTL 自動刪除舊訊息

---

## 🔸 2. 如何設計一個遊戲紀錄系統，每日千萬筆寫入？

### 🎯 考點
- 水平擴充（Sharding）
- 寫入策略、partition key 選擇
- 熱點分散與寫入順序

### 💡 建議方向

- 以 `userId` + `timestamp` 做 shard key 分散
- 避免 `_id` 作 shard key（單調遞增會導致單一 shard）
- 按月/週分表可緩解單 collection 壓力

---

## 🔸 3. 如何設計強一致性的查帳系統？

### 🎯 考點
- WriteConcern / ReadConcern
- Transaction 運用時機
- 錯誤恢復策略

### 💡 建議方向

- 使用 `writeConcern: majority + j:true` 確保資料複寫成功
- 查詢時使用 `readConcern: majority`
- 多集合可考慮 multi-document transaction（需要 replica set）

---

## 🔸 4. 在什麼情況下你會選擇 MongoDB 而不是 MySQL？

### 🎯 考點
- 非結構化資料
- 快速開發、需求變動頻繁
- 多層巢狀資料

### 💡 建議方向

| 使用 MongoDB 的場景                   |
|---------------------------------------|
| 活動紀錄 / 使用者事件流               |
| 即時聊天 / 訊息串                      |
| IoT 裝置資料 / 非同步設備訊息收集     |
| Schema 經常變動 / rapid iteration     |
| 需要地理位置索引（GeoJSON）           |

---

## 🔸 5. MongoDB 的 B-Tree index 有什麼設計限制？

### 🎯 考點
- 單值 vs 陣列 vs 多欄位索引
- 過大 document 或欄位分布不均時效能影響
- 預熱與覆蓋查詢（covered index）

---

## 🔸 6. 實務上如何應對 Mongo 寫入壓力過大？

### 🎯 建議解法

- 前端透過 Kafka queue 緩衝
- 批次寫入（BulkWrite）
- 設定 shard + 合理 shard key（避免寫入單一分片）
- 分表（按日/月）搭配 TTL

---

## 🔸 7. 如何處理 Mongo 的 retry / failover？

### 🎯 關鍵設定

- `retryWrites: true`
- `retryReads: true`
- 瞭解 failover 後 driver 的 reconnect 行為
- 保守設計 application 層級的 retry with backoff

---

## 🔸 8. 一筆資料是否適合用 embedding？

### ✅ 適合的情境

| 條件                       | 建議 |
|----------------------------|------|
| 子文件數量少（< 100）      | 👍   |
| 子資料無需頻繁更新         | 👍   |
| 每次都會一起查詢或顯示     | 👍   |

### ❌ 不適合的情境

| 條件                          | 建議 |
|-------------------------------|------|
| 子文件數量會超過 1MB 文件限制 | ❌   |
| 子資料需獨立更新/查詢頻繁    | ❌   |

---

## 📌 面試補充 Tips

- 瞭解 Mongo 的強項在於彈性 schema 與讀取模式彈性
- Sharding 不是萬靈丹，須在吞吐達瓶頸才評估
- 注意 TTL、index size、16MB document 限制等底層限制
