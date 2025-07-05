# 10. Monitoring and Profiling

本章節介紹如何在 MongoDB 中進行監控與效能分析，對於架構設計、異常排查與容量預測至關重要。對於 Staff Engineer，需要能針對問題快速定位與觀察系統瓶頸。

---

## 📊 一、MongoDB 內建監控工具

### 1. `serverStatus`
提供整體執行狀況，如記憶體、索引命中率、寫入延遲。

```js
db.serverStatus()
```

常用觀察項目：

| 指標               | 意義                           |
|--------------------|--------------------------------|
| `connections`      | 當前連線數、可用連線數         |
| `wiredTiger.cache` | 記憶體快取命中率               |
| `opcounters`       | CRUD 操作統計                  |
| `locks`            | 資源鎖情況（可檢測瓶頸）       |

---

### 2. `currentOp` 與 `killOp`

- 查詢目前進行中的操作（包含慢查詢）
```js
db.currentOp({ "secs_running": { "$gte": 2 } })
```

- 手動終止特定操作
```js
db.killOp(<opid>)
```

---

### 3. `collStats` / `dbStats`

- 檢視 collection 大小、索引數量與碎片情況
```js
db.collection.stats()
```

---

## 🐢 二、慢查詢分析與 profiler

### 1. 啟用 profiler

```js
db.setProfilingLevel(1, { slowms: 100 })
```

- Level 0：不記錄
- Level 1：記錄慢查詢（預設 > 100ms）
- Level 2：記錄所有查詢（高風險）

### 2. 查詢 profiler 資料

```js
db.system.profile.find().sort({ ts: -1 }).limit(5)
```

- 可觀察慢查詢的索引使用、執行時間、query plan 等

---

## 📈 三、監控工具

| 工具            | 說明                                    |
|-----------------|-----------------------------------------|
| **MongoDB Atlas** | 雲端 UI 提供即時指標與警示              |
| **mongostat**     | CLI 工具，提供即時操作統計               |
| **mongotop**      | 類似 top，可顯示各 collection IO 活動    |
| **Prometheus Exporter** | 適合自建監控平台結合 Grafana        |

---

## 🧠 Staff Engineer 該理解的點

- 了解如何快速定位慢查詢瓶頸（`explain`, `profile`, `collStats`）
- 定期監控系統指標變化，提早預測資源不足
- 熟悉慢查詢修正手法（加索引、refactor query、資料重整）
- 搭配 CI/CD，自動化觀察查詢變化與異常警示

---

[← 回到總覽](../Mongo_Summary.md)