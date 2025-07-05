# 01. Mongo Overview

本章節介紹 MongoDB 的核心概念、架構組成與它與傳統關聯式資料庫的差異。作為 Staff Engineer，理解 MongoDB 的整體運作原理是後續進行效能分析、設計 tradeoff 與系統選型的基礎。

---

## 🔧 MongoDB 是什麼？

MongoDB 是一個開源的 NoSQL 文件導向資料庫，具備以下特性：

- 使用 **BSON（Binary JSON）** 格式儲存資料
- 支援動態 schema（不需事先定義欄位）
- 支援高可用（Replica Set）、分散式（Sharding）架構
- 提供強大的查詢語言與 aggregation pipeline

---

## 📁 基本結構

| 概念層級     | 描述                                  |
|--------------|---------------------------------------|
| **Database** | 多個 collection 的集合                 |
| **Collection** | 類似 RDBMS 的 table                  |
| **Document** | 類似 RDBMS 的 row，格式為 BSON        |
| **Field**     | 類似 RDBMS 的 column（支援嵌套）      |

範例 document：
```json
{
  "_id": ObjectId("..."),
  "name": "Alice",
  "age": 30,
  "tags": ["admin", "premium"],
  "profile": {
    "email": "alice@example.com",
    "joined": "2022-01-01"
  }
}
```

---

## 🔂 與關聯式資料庫的差異

| 比較項目 | MongoDB                            | MySQL/PostgreSQL                    |
|----------|-------------------------------------|-------------------------------------|
| 資料格式 | BSON (Document)                    | Row-based (table schema)            |
| Schema   | 動態（可每筆不同）                 | 固定欄位結構                         |
| Join 支援 | 限制（$lookup）                   | 多表 join 靈活                       |
| Transaction | 支援（ReplicaSet / Shard 有限制） | 原生強一致交易                       |
| 水平擴充 | 原生支援 Sharding                  | 需手動切分                           |

---

## 🧱 MongoDB 架構角色

| 元件         | 功能說明                            |
|--------------|-------------------------------------|
| `mongod`     | 資料節點，負責儲存與處理資料         |
| `mongos`     | 分片環境下的 query 路由器            |
| `config server` | 儲存 metadata（分片、chunk、cluster 狀態） |

MongoDB 可部署成：
- 單機（單一 mongod）
- Replica Set（高可用，主從複製）
- Sharded Cluster（水平分片擴充）

---

## 🧪 運作流程簡圖

```
Client
  ↓
[ Driver / mongos ]
  ↓
[ mongod (Primary) ] ←→ [ mongod (Secondary) ]
```

---

## 🔍 補充重點

- MongoDB 的 `_id` 欄位預設為主鍵，資料寫入時必須唯一。
- BSON 格式支援比 JSON 更多的資料型別，例如 Date、Binary、Decimal128。
- 支援內建 replication、sharding、TTL、全文索引等功能。

---

## 🧠 Staff Engineer 該理解的點

- MongoDB 適用於 schema 靈活、多層巢狀、橫向擴充的場景（如：遊戲紀錄、event log、聊天紀錄）
- 了解它的資料模型限制（如 document 最大 16MB）與擴充模式
- 熟悉其架構是後續設計 shard key、replica failover、交易策略的基礎

---

[← 回到總覽](../Mongo_Summary.md)
