# 09. Performance and Tuning

本章節探討 MongoDB 在實務上進行效能優化與參數調整的建議，對於處理高吞吐寫入、大量查詢、批次處理等需求場景尤為重要。

---

## 🧠 一、Schema 設計優化

- **避免深層嵌套 document**
  - MongoDB 預設最大 depth 為 100 層，避免不必要巢狀導致效率下降
- **控制 document 大小**
  - 單筆 document 建議控制在幾 KB 內，最多 16MB
- **選擇適當資料類型**
  - `NumberLong` vs `int`，節省空間可提升記憶體命中率

---

## ⚡ 二、索引（Index）優化技巧

- 使用 **compound index** 對應複合查詢欄位
- 索引設計要依照實際 query pattern（WHERE, SORT）
- **避免覆蓋 index 冗餘欄位**
- 使用 `hint()` 測試與強制使用某 index

---

## 🔍 三、查詢效能提升建議

- 查詢應搭配索引欄位，避免 collection scan
- 控制傳回欄位，使用 `projection`
- 使用 `explain("executionStats")` 分析查詢瓶頸
- 善用 `readConcern` 搭配延遲容忍範圍

---

## 🧵 四、寫入效能優化技巧

- 使用 **bulkWrite()** 進行大量寫入
- 降低寫入確認強度（如：`w:1`、不等 `majority`）
- 搭配 **WriteConcern + Journal 設定** 調整安全性與效能
- 調整 `_id` 輸入順序（可考慮非 ObjectId）

---

## 🏷 五、Aggregation Pipeline 最佳化

- 節省記憶體運算量，將 `$match`、`$project` 放前面
- 避免 `$unwind` + `$group` 導致大量 interim result
- 使用 `$merge` 輸出結果避免返回過多記憶體資料
- 加上索引支援 pipeline 首段查詢

---

## 🔧 六、系統參數調整

- 調整 WiredTiger cacheSize（約為實體記憶體 50%）
- 對應 SSD 寫入優化，調整 journal commitInterval
- 網路設定如 `maxIncomingConnections`、`backlog`

---

## 📈 七、觀察效能的方式

- 使用 `serverStatus` 觀察 cache hit rate、locks 等
- `db.currentOp()` 檢查長時間操作與慢查詢
- 搭配 MongoDB Atlas / Prometheus / Cloud Monitoring

---

## 🧠 Staff Engineer 該理解的重點

| 面向           | 需掌握的概念                                       |
|----------------|----------------------------------------------------|
| 資料設計       | schema 大小、索引規劃、query pattern 分析         |
| 寫入效能       | bulkWrite、writeConcern 調整、安全性 vs 效能平衡 |
| 查詢優化       | explain 分析、filter 與投影、aggregation 規則     |
| 系統調校       | 儲存引擎參數、記憶體快取策略、硬體資源評估         |

---

[← 回到總覽](../Mongo_Summary.md)