# 04. Read & Write Concern

MongoDB 提供讀寫時的一致性與耐久性控制機制，分別為 **Read Concern** 和 **Write Concern**。這些設計允許開發者針對操作需求選擇最適合的強度平衡 CAP。

---

## 📝 Write Concern

Write Concern 控制「資料寫入操作的確認程度」。

| 層級     | 描述                                             |
|----------|--------------------------------------------------|
| `w: 0`   | 不確認寫入結果（fire-and-forget）                |
| `w: 1`   | 寫入 Primary 即返回確認                          |
| `w: "majority"` | 寫入被多數節點確認後才回應                     |
| `j: true`| 要求操作寫入 journal，確保掉電也不會遺失資料         |
| `wtimeout`| 設定等待確認的最大時間（毫秒）                    |

### 實例：
```js
db.users.insertOne(
  { name: "Alice" },
  { writeConcern: { w: "majority", j: true, wtimeout: 1000 } }
)
```

---

## 🔍 Read Concern

Read Concern 控制「讀取資料的可見性與一致性層級」。

| 層級             | 描述                                                            |
|------------------|-----------------------------------------------------------------|
| `local`（預設）   | 讀取 Primary 的本地資料，無法保證同步                          |
| `available`      | 允許讀取 Secondary，可能有 lag 或不一致                         |
| `majority`       | 讀取大多數節點確認的資料（確保資料被多數節點寫入）              |
| `linearizable`   | 最強一致性，保證 Primary 最新、序列化（效能最差）               |
| `snapshot`       | 針對交易讀取的快照一致性，僅限於 multi-document transaction 使用 |

### 實例：
```js
db.users.find({}, { readConcern: { level: "majority" } })
```

---

## 🧠 Read / Write Concern 實務應用建議

| 場景               | 建議設計                       |
|--------------------|--------------------------------|
| 即時聊天系統         | `w: 1` + `readConcern: local` |
| 金流交易寫入         | `w: "majority", j: true`       |
| 玩家歷史查詢         | `readConcern: majority`        |
| 高併發批次寫入       | `w: 1`（搭配非同步批次寫）     |

---

## ⚠️ 注意事項

- **Read Concern 只保證讀取一致性，不保證寫入成功**
- 若 Secondary 同步落後，`readConcern: majority` 可能會讀不到最新資料
- Linearizable 模式效能極差，應慎用

---

## 🧠 面試重點提示（Staff Engineer）

- Mongo 提供 **調整一致性與效能的彈性機制**
- 實務上選擇關鍵在於場景要求（高效能 vs 高一致性）
- 熟悉各層級意義與搭配是設計可靠分散式系統的基礎

---

[← 回到總覽](../Mongo_Summary.md)