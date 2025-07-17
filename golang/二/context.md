## Go `context` 面試全攻略筆記（適用於 Staff Engineer）

---

### ✅ 背後設計目的

| 項目 | 說明 |
|------|------|
| 核心用途 | 解決跨 API 的取消控制、逾時處理、請求鏈追蹤（不是資料交換） |
| 設計哲學 | 避免 goroutine 泄漏，確保資源釋放與控制流一致 |

---

### ✅ 基本 API 與使用情境

```go
ctx := context.Background()
ctx, cancel := context.WithCancel(ctx)
defer cancel()

go func(ctx context.Context) {
    select {
    case <-ctx.Done():
        // 中止 goroutine
    }
}(ctx)
```

| API | 功能 |
|-----|------|
| `context.Background()` | root context，用於 main、init、test 開始點 |
| `context.TODO()` | 尚未決定用法的 placeholder context |
| `context.WithCancel(parent)` | 可呼叫 cancel() 結束下游 context |
| `context.WithTimeout(parent, duration)` | 指定逾時時間，自動 cancel |
| `context.WithDeadline(parent, time.Time)` | 指定固定時間點 deadline |
| `context.WithValue(parent, key, val)` | 搭配 key-value 傳遞 metadata（**避免業務邏輯使用**） |

---

### ✅ `WithValue` 使用建議

- ✅ 用於：trace ID、request ID、log context、user info
- ❌ 避免：業務參數、邏輯流程控制（會造成耦合與可測試性低）

```go
ctx := context.WithValue(ctx, userIDKey, 12345)
```

---

### ✅ 搭配 select 控制 goroutine 終止

```go
select {
case <-ctx.Done():
    log.Println("context cancelled")
case result := <-work:
    return result
}
```

---

### ✅ 鏈式設計實例

```go
ctx, cancel := context.WithCancel(context.Background())
ctx, timeoutCancel := context.WithTimeout(ctx, 2*time.Second)
```

---

### ✅ 常見錯誤與陷阱

| 錯誤行為 | 為什麼錯？ |
|-----------|-------------|
| 將 context 傳給未知的 goroutine | 可能會失去控制，無法 cancel |
| 在 handler 內建構全新的 context | 不可觀測 / 無法追蹤父流程 |
| 用 WithValue 傳遞業務邏輯 | 測試困難、耦合度高 |
| 不呼叫 cancel | 造成 goroutine 泄漏，或資源佔用 |

---

### ✅ 實戰場景

#### HTTP / gRPC
- 請求流程控制（cancel 時不繼續處理）
- 超時設計
- 傳遞 traceID

#### 資料庫操作
- 控制 query timeout / cancel query
- 清理連線與 context-aware 驅動整合

#### 併發任務
- 等待任務完成或逾時（`select { <-ctx.Done() }`）

#### 微服務追蹤
- 透過 `context.WithValue` 傳遞 trace ID / span ID
- 將 log 訊息綁定 context log field

---

### ✅ 延伸面試問題（Staff Engineer）

| 問題 | 回答方向 |
|------|----------|
| context 的目的？ | 控制 cancel / timeout / 請求鏈追蹤 |
| 如何實作 timeout / cancel？ | `WithTimeout`, `WithCancel`, `<-ctx.Done()` |
| context 與 goroutine 的正確搭配？ | select + Done, 並確保 cancel 被呼叫 |
| 為何 context 是 interface？ | 易於擴展與實作遞迴式 cancel 模型 |
| 哪些情況不能用 WithValue？ | 不可作為業務資料傳遞機制 |

---

### ✅ 小結

- `context` 是 Go 架構中的關鍵設計，特別在高併發、微服務、gRPC 等場景
- API 雖少，但內含豐富設計哲學
- 不只是會用，更要懂設計背後的原則與限制

---

如需更多例子，可補充 goroutine 超時管理 / context tree 結構圖。