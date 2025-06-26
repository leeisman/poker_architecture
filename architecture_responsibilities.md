# 🎮 遊戲服務職責說明文件

本文件列出系統中兩大核心模組：`room_server` 與 `table_server` 的設計職責與邊界區分，協助團隊理解模組責任、維運範疇與系統架構的清晰性。

---

## 🧭 room_server — 房間與金流管理核心

### 🎯 定位與責任
room_server 為遊戲中介與調度層，專責處理玩家進出桌、金流整合、比賽控制等任務，並為 table_server 提供狀態支援與控制指令。

### 📌 核心職責

| 模組             | 說明 |
|------------------|------|
| 💰 金流處理      | 負責處理 `rebuyREQ`、`addonREQ` 等指令，透過 ValueServer 完成金流交易 |
| 🧭 玩家調度      | 管理 `enterRoomREQ` / `leaveRoomREQ`，負責資金鎖定、轉帳與進出紀錄 |
| 🪑 配桌管理      | 控制建桌邏輯（含 least_conn），並將玩家分配至適當 table_server |
| 🧾 狀態同步快取  | 從 Mongo 讀取牌桌當前狀態，例如晉桌人數等等 |
| 🔁 MTT 賽制控制  | 控制 MTT 桌次生成、淘汰者處理、合桌、進度同步等流程 |
| 🏁 結算發獎      | 統整名次與獎金分配，調用金流服務完成最終發獎 |

### ❌ 不負責的部分

- 遊戲邏輯與玩法流程（交由 table_server）
- WebSocket 長連線處理（由 Gateway 處理）
- 牌局紀錄與 DB 寫入（由 game_record_server 負責）

---

## 🂡 table_server — 遊戲玩法邏輯執行核心

### 🎯 定位與責任
table_server 為每一個「牌桌」的運作主體，負責所有牌局的玩法處理、玩家行為與流程控制。

### 📌 核心職責

| 模組              | 說明 |
|-------------------|------|
| 🃏 遊戲流程控制    | 管理發牌、下注、比牌等核心邏輯，每桌為獨立狀態機 |
| 💬 玩家行為處理    | 處理如 `iBetREQ`、`iFoldREQ` 等所有 gameplay 指令 |
| ⏱️ Timeout 控制    | 透過狀態機追蹤玩家動作，雙回合未動作即自動 fold 或移除 |
| 🧾 結束回報        | 每回合結束後產生 `hand_record` 上報 game_record_server |
| 🔁 合桌支援        | 接收 `iMigratePlayers` 指令，負責玩家座位安排與同步進桌資訊 |
| 📤 Kafka 發送      | 將 `poker.hand.result` 透過grpc請求到 game_record_server 發送至 Kafka 供活動模組等訂閱處理 |

### ❌ 不負責的部分

- 金流處理與轉帳（由 room_server 控管）
- 玩家是否允許進桌（room_server 驗證後指派）
- 直接資料庫存取（所有存取統一透過 game_record_server）
- 玩家快取與認證驗證（由 Gateway/Auth 處理）

---

## ✅ 協作原則備註

- room_server 為 **金流與配桌調度中心**，table_server 為 **純遊戲邏輯執行體**。
- 所有需金錢操作的流程，如 `rebuy`, `addon`, `leaveRoom` 都應透過 room_server。
- 純牌桌遊戲行為如 `Bet`, `Fold`, `Check` 等則由 table_server 直接處理。
- 兩者間 **避免互相呼叫**，room_server 可下指令至 table_server，但反向不行。
- table_server 所有寫入請交由 game_record_server 轉寫 Mongo，或發送 Kafka 事件供異步處理。

---