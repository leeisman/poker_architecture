# 03. Sharding and Distribution

本章節將說明 MongoDB 的 Sharding 架構設計，包括 Shard、Chunk、Shard Key 等基本觀念，並探討如何分散儲存與處理負載以支援大規模資料與高可用架構。

---

## 🧱 Sharding 是什麼？

Sharding 是 MongoDB 實現 **水平擴充（horizontal scaling）** 的方法。

- 將資料依據某個欄位（Shard Key）切分成多個 Chunk
- 將 Chunk 分散儲存在多個 Shard 節點上
- `mongos` 為查詢路由器，負責將請求導向正確的 Shard

---

## 📦 Sharding 架構圖

```
Client
  ↓
mongos (Query Router)
  ↓
+---------+     +---------+     +---------+
| Shard 1 |     | Shard 2 |     | Shard 3 |
| P + S1  |     | P + S1  |     | P + S1  |
+---------+     +---------+     +---------+
       ↖          ↖           ↖
      config server (metadata 儲存)
```

---

## 🧩 架構元件說明

| 元件             | 說明                                             |
|------------------|--------------------------------------------------|
| `Shard`          | 一個 Shard 可是 Replica Set（有 Primary / Secondary） |
| `mongos`         | 路由器，處理 Client 查詢，並分發給正確 Shard        |
| `Config Server`  | 儲存所有 Chunk metadata、分片資訊、路由規則         |
| `Chunk`          | 實際被切割與分配到各個 Shard 的資料範圍              |

---

## 🗝️ Shard Key 是什麼？

Shard Key 決定資料如何被切分與放置在各個 Shard。

| 設定錯誤風險 | 範例                                  |
|--------------|---------------------------------------|
| 熱點集中     | 使用 `_id` 或遞增數字會導致寫入集中於單一 Shard |
| 分散不均     | 使用太多重複值欄位（如 status）造成不平均分佈   |

🧠 Shard Key 一旦建立後 **無法更改**，需謹慎設計！

---

## 🧮 Chunk Migration（動態平衡）

- MongoDB 會自動偵測某個 Shard 負載過重
- 透過 Chunk Migration 將部分 Chunk 移到其他 Shard
- 避免單點過載，提高分散效能

---

## 🧠 Staff Engineer 應理解的重點

- Sharding 提供 MongoDB 真正的「無限擴充」能力
- 設計良好的 Shard Key 是分散讀寫與避免熱點的關鍵
- `mongos` 與 Config Server 是架構瓶頸／可用性關鍵
- 搭配 Replica Set，才能做到**橫向擴充 + 高可用**

---

[← 回到總覽](../Mongo_Summary.md)
