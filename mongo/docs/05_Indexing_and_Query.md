# 05. Indexing and Query

本章節說明 MongoDB 的索引結構、查詢行為與效能關聯。作為 Staff Engineer，掌握索引設計的策略是系統效能優化的關鍵之一。

---

## 🔍 為什麼需要 Index？

- 沒有索引時，查詢會進行 **collection scan**
- 加入索引可將查詢從 O(n) 降為 O(log n)
- 支援高效的範圍查詢、排序、唯一性驗證等

---

## 🧱 索引種類

| 類型             | 說明                                           |
|------------------|------------------------------------------------|
| 單欄位索引       | 最常見，用於常查詢欄位                        |
| 複合索引         | 針對多欄位設計查詢順序，左側欄位必須命中      |
| 唯一索引         | 強制欄位值不能重複                             |
| TTL 索引         | 設定自動過期時間（常用於 log / session）     |
| Hashed 索引      | 用於 Sharding 的 shard key                    |
| Text 索引        | 支援全文搜尋（不支援複合索引）               |
| Geospatial 索引  | 地理空間搜尋，如 `$near`, `$geoWithin`        |

---

## 🔧 建立與使用索引

```js
db.users.createIndex({ age: 1 }) // 遞增索引
db.users.createIndex({ name: 1, age: -1 }) // 複合索引
db.users.find({ age: { $gt: 18 } }).sort({ name: 1 })
```

可使用 `explain()` 查看查詢計畫：

```js
db.users.find({ age: { $gt: 18 } }).explain("executionStats")
```

---

## ⚠️ 索引設計注意事項

- 不必要的索引會降低寫入效能（每次寫入需更新索引）
- 複合索引需命中「最左前綴原則」
- 過多索引會消耗大量記憶體
- Text 與 Geospatial 索引無法組合使用

---

## 📈 範例索引命中分析

```js
db.orders.createIndex({ user_id: 1, created_at: -1 })

// 命中索引 ✅
db.orders.find({ user_id: 123 }).sort({ created_at: -1 })

// 不命中索引 ❌（user_id 沒出現在查詢條件中）
db.orders.find({ created_at: { $gt: ISODate(...) } })
```

---

## 🧠 Staff Engineer 該理解的點

- 避免 over-index（索引過多）
- 應依據查詢 pattern 設計索引結構
- 熟悉 `explain()` 工具來分析查詢效能瓶頸
- 了解 shard key 也需要設計成索引

---

[← 回到總覽](../Mongo_Summary.md)