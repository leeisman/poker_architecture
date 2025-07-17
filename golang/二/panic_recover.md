
## Go 面試筆記：panic / recover（Staff Engineer 適用）

---

### ✅ 面試常見問題

| 問題               | 答案 |
|--------------------|------|
| 如何捕捉 panic？   | 使用 `defer + recover()` 包住可能觸發 panic 的邏輯 |
| 應用在哪些場景？   | goroutine crash 保護、server handler 安全保底、actor framework 防止整體掛掉 |

---

### ✅ 實戰原則

- `recover()` 僅能在 `defer` 函式中使用
- panic 發生後程式會停止執行，進入 defer stack
- 只能 recover 當下 goroutine 的 panic，無法捕捉其他 goroutine 的 panic

---

### ✅ 使用時機與陷阱

| 項目 | 建議 |
|------|------|
| ✅ 適用場景 | goroutine crash 保護、middleware 全域保底、防止整體服務崩潰 |
| ❌ 不推薦 | 當作錯誤處理機制、取代正常的 error return |
| 陷阱 | recover 無效時機（非 defer 中）、不 log 導致 debug 困難 |

---

### ✅ 範例程式碼

```go
func safeExecute() {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("panic caught: %v", r)
        }
    }()
    dangerousWork()
}
```

---

### ✅ 延伸應用

1. **HTTP Middleware 中間層保護：**

```go
func recoverMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Println("Recovered from panic:", err)
                http.Error(w, "Internal Server Error", 500)
            }
        }()
        next.ServeHTTP(w, r)
    })
}
```

2. **Worker Pool / goroutine crash 捕捉：**

```go
go func() {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("worker panic: %v", r)
        }
    }()
    doWork()
}()
```

---

### ✅ 面試深入問法

| 問題 | 應對方向 |
|------|----------|
| recover 能不能跨 goroutine？ | 不能，只能作用於當前 goroutine 的 defer stack |
| defer/recover 效能成本？ | defer 有輕微開銷，panic 時成本大（因為需建構 stack trace） |
| recover 如何搭配 logging / metrics？ | 通常與 log 結合記錄錯誤，並可能上報給 APM 工具（如 Sentry, Datadog） |

---

### ✅ 小結

- panic/recover 是保底設計而非錯誤處理主流程
- 在高併發 / 遊戲服務中非常實用，用來防止 crash 拉垮整體服務
- 搭配 defer 形成安全邊界，與 context / log 搭配最佳化服務穩定性
