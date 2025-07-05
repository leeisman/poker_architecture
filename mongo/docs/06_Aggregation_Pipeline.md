# 06. Aggregation Pipeline

Aggregation Pipeline 是 MongoDB 提供的資料處理框架，類似 SQL 的 GROUP BY、JOIN、ORDER BY 等複雜查詢操作，透過階段性處理的方式操作文件集合。

---

## 🔧 什麼是 Aggregation Pipeline？

MongoDB Aggregation 是一種資料處理方式，它允許將多個「處理階段」串接起來，每個階段使用一個操作符（如 `$match`、`$group`）來逐步轉換資料。

### 基本語法：

```js
db.collection.aggregate([
  { $match: { status: "active" } },
  { $group: { _id: "$user_id", total: { $sum: "$amount" } } },
  { $sort: { total: -1 } }
])
```

---

## 🔄 Pipeline 階段介紹

| 階段       | 說明                                               |
|------------|----------------------------------------------------|
| `$match`   | 篩選條件，功能類似 `WHERE`                         |
| `$group`   | 群組資料，功能類似 `GROUP BY`                      |
| `$project` | 篩選/轉換欄位，功能類似 `SELECT`                   |
| `$sort`    | 排序資料                                           |
| `$limit`   | 限制回傳筆數                                       |
| `$skip`    | 跳過前 N 筆資料                                    |
| `$lookup`  | 類似 JOIN，可跨 collection 關聯                    |
| `$unwind`  | 將陣列欄位展開成多筆文件                           |
| `$addFields` | 新增欄位                                          |

---

## 📦 範例：計算每位玩家總下注金額

```js
db.game_logs.aggregate([
  { $match: { event: "bet" } },
  { $group: { _id: "$player_id", total_bet: { $sum: "$amount" } } },
  { $sort: { total_bet: -1 } }
])
```

---

## 🔍 範例：查詢最近 7 天內的所有行為並依玩家分群

```js
db.actions.aggregate([
  { $match: { timestamp: { $gte: ISODate("2025-07-01") } } },
  { $group: {
      _id: "$user_id",
      actions: { $push: "$action_type" }
  }}
])
```

---

## 🧠 Staff Engineer 該理解的重點

- Pipeline 是 MongoDB 中最有彈性與威力的查詢與轉換工具。
- 必須理解各階段執行順序與執行計劃，才能優化效能。
- 注意 `$group`、`$sort` 常會觸發 memory 限制（預設 100MB）。
- 使用 `$merge`、`$out` 寫回 collection 時應考慮併發與交易需求。

---

## 📈 效能調校建議

- 優先使用 `$match` 篩選資料，避免不必要的運算。
- 配合適當索引讓 `$match`、`$sort` 發揮作用。
- 可使用 `allowDiskUse: true` 避免記憶體限制錯誤。

---

[← 回到總覽](../Mongo_Summary.md)
