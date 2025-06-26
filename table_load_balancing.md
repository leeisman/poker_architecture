## 🔀 3. Load Balancing（負載平衡）
---

```mermaid
sequenceDiagram
    participant Room as room_server
    participant Mongo as MongoDB
    participant Table1 as table_server_1
    participant Table2 as table_server_2

    Room->>Mongo: aggregate table_server_id + conn_count
    Mongo-->>Room: 回傳 table_server_1:2, table_server_2:5
    Room->>Table1: 建立新桌子 table_abc
```