# 08. Schema Design and Modeling

MongoDB 提供靈活的 schema 設計，但靈活性同時也帶來設計挑戰。本章重點在於如何選擇嵌套（Embedding）或關聯（Referencing），以及 Schema 如何影響效能與擴充性。

---

## 🧱 基本設計單位：Document

MongoDB 的基本儲存單位是 document（BSON 物件），每筆 document 最大為 16MB，支援巢狀與陣列結構。

---

## 📦 Embedding vs. Referencing

| 類型         | 優點                                      | 缺點                                   |
|--------------|-------------------------------------------|----------------------------------------|
| **嵌入**     | 快速查詢（單一讀取）                     | 文件大小受限、不易獨立查詢             |
| **參考**     | 模組化設計、便於擴充                     | 查詢需多次查詢或 `$lookup`              |

### 範例：

#### 嵌入（Embedding）

```json
{
  "user_id": 1,
  "name": "Alice",
  "address": {
    "city": "Taipei",
    "zip": "100"
  }
}
```

#### 參考（Referencing）

```json
// users collection
{ "_id": 1, "name": "Alice" }

// addresses collection
{ "user_id": 1, "city": "Taipei", "zip": "100" }
```

---

## 🏗️ Schema 設計原則

1. **根據讀取頻率設計**：常用資料應靠近一起（嵌入）。
2. **避免超過 document 限制（16MB）**：大型清單要分開。
3. **資料寫入頻繁者避免嵌入**：避免整包資料重新寫入。
4. **考慮 transaction**：跨多 collection 需考慮交易支援。
5. **版本控管**：可在 document 中加入 `_v` 欄位管理格式版本。

---

## 🧠 使用場景比較

| 使用情境                       | 建議方式     |
|------------------------------|--------------|
| 聊天紀錄、留言                | Referencing  |
| 使用者設定檔                  | Embedding    |
| 商品與規格                    | 視讀取與更新頻率決定 |
| 手遊活動紀錄（game logs）     | Referencing（避免文件過大）|

---

## 🔄 Schema 動態變化

MongoDB 支援動態欄位，不需事前定義 schema。但：

- 仍建議在應用程式層定義資料模型（ex: Go struct）
- 建立索引與效能分析仍依賴穩定的欄位設計

---

## 🧠 Staff Engineer 該理解的點

- 避免「把關聯資料全部丟一起」造成 bloated document
- 針對不同服務（如 Game Log vs. Player Profile）做資料模型拆分
- 嵌入與參考設計應與資料熱度、更新頻率及一致性需求綜合評估

---

## 📚 延伸閱讀

- [MongoDB 官方 schema 設計模式](https://www.mongodb.com/blog/post/6-rules-of-thumb-for-mongodb-schema-design-part-1)

---

[← 回到總覽](../Mongo_Summary.md)
