# 🎰 Jackpot 累積與發獎鎖設計說明

本設計目的是處理高併發下注場景下的 Jackpot 累積與中獎結算流程，避免資料錯誤或發獎重複。透過 Redis 實作兩層鎖機制達到正確性與效率平衡。

---

## 🔐 鎖的角色與責任

| 鎖名稱                     | 類型                 | 作用                                  | 使用時機            |
|----------------------------|----------------------|---------------------------------------|---------------------|
| `reward_pool_lock`         | Redis 分布式鎖       | 保護獎池累積金額 + 寫入紀錄           | 每次下注累積 Jackpot |
| `priority_reward_pool_lock`| Redis key + 分布式鎖 | 表示當前正在發獎，其他流程不可進入     | 玩家中獎觸發結算時   |

---

## 📌 核心邏輯流程

### 累積獎池流程（下注時）：

```go
// 檢查是否正在發獎
if redis.GET(priority_reward_pool_lock_key) != nil {
    // 有人在發獎，暫停後重試
    sleepForPriorityRewardPoolLock()
    retry FeeLock()
}

// 若未鎖，則嘗試鎖住 reward pool 進行加總
lock(reward_pool_lock)
累加 Jackpot 金額
寫入紀錄
unlock(reward_pool_lock)
```
---
## 發獎流程（中獎時）：

```go
// 嘗試設定 priority lock，表示發獎進行中
SET priority_reward_pool_lock_key "1" EX 10 NX
若 SET 失敗，代表已有其他人發獎，應等候或退出

lock(reward_pool_lock)
讀取並清空 jackpot
發獎並寫入紀錄
unlock(reward_pool_lock)

DEL priority_reward_pool_lock_key // 解鎖發獎狀態
```
---
## 上面做法累積獎池流程還是會搶到，需要設計double-check
```go
// 嘗試進行獎池累積（Double-Check 保護）
for {
    // Step 1: 檢查是否正在發獎
    if redis.GET(priority_reward_pool_lock_key) != nil {
        sleepForPriorityRewardPoolLock()
        continue
    }

    // Step 2: 嘗試鎖住 reward pool
    lock(reward_pool_lock)

    // Step 3: 再次檢查是否正在發獎（double-check）
    if redis.GET(priority_reward_pool_lock_key) != nil {
        unlock(reward_pool_lock)
        sleepForPriorityRewardPoolLock()
        continue
    }

    // Step 4: 安全執行累積與紀錄
    累加 Jackpot 金額
    寫入紀錄
    unlock(reward_pool_lock)
    break
}
```