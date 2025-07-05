# 🔁 Table Migration 桌子遷移流程設計

當進行 **滾動部署** 或遇到 table_server 發生異常時，room_server 需主動觸發桌子遷移機制，以確保遊戲不中斷並支援狀態復原。以下為完整的流程與設計說明。

---

## 🎯 遷移目標

- ✅ 無縫切換 Table 實例，不中斷玩家遊戲
- ✅ 保留當前遊戲狀態並還原（含 queue 中未處理指令）
- ✅ 自動更新 GameRouter 的 routing 資訊
- ✅ 支援原 Table 緩衝 → 新 Table 消化的轉接機制

---

## 🧭 遷移流程圖

```mermaid
sequenceDiagram
    participant Room as room_server
    participant OldTable as 舊 table_server
    participant Redis as Redis
    participant NewTable as 新 table_server
    participant GameRouter as GameRouter (指令快取路由)

    %% Step 1: 偵測需遷移桌子
    Room->>OldTable: HealthCheck / 停止接收新請求

    %% Step 2: 舊 table_server 將狀態寫入 Redis
    OldTable->>Redis: SET table_state:table_123 = {...} EX 300

    %% Step 3: 舊 table_server 仍接收指令，寫入 queue
    Note over OldTable: 將指令寫入 Redis queue
    OldTable->>Redis: LPUSH table_queue:table_123 BetREQ(...)

    %% Step 4: 啟動新 table_server 做復原
    Room->>NewTable: recovery_table(table_123)

    %% Step 5: 新 table_server 載入狀態
    NewTable->>Redis: GET table_state:table_123

    %% Step 6: 宣告上線 + 更新心跳
    NewTable->>Redis: SET table_alive:table_123 = "new_instance" EX 90

    %% Step 7: 消化 queue 直到遇到 end
    loop 消化 Redis queue
        NewTable->>Redis: RPOP table_queue:table_123
        alt 指令為 end
            Note right of NewTable: 停止消化，遷移結束
        else 正常指令
            NewTable->>NewTable: 處理 BetREQ(...)
        end
    end

    %% Step 8: 遷移完成通知
    NewTable-->>Room: Migration Ready

    %% Step 9: 通知 OldTable 關閉
    Room->>OldTable: MigrationEnd(table_123)

    %% Step 10: OldTable 寫入 end
    OldTable->>Redis: LPUSH table_queue:table_123 end

    %% Step 11: 更新 GameRouter 快取
    Room->>GameRouter: Update table routing (table_123 → new_instance)
```