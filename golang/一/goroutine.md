# 🧵 Golang Goroutine 重點整理

---

## 🧠 Goroutine 是什麼？

- Golang 中的輕量級執行緒，由 Go runtime 管理（不是 OS thread）
- 透過 `go` 關鍵字啟動，例如：`go doSomething()`
- 可同時啟動成千上萬個 goroutines，記憶體開銷極低（初始只佔 2KB stack）

---

## 🔄 GMP 調度模型（重中之重）

| 元件 | 說明 |
|------|------|
| G (Goroutine) | 執行單元，每個 go func 都是一個 G |
| M (Machine)   | 真正的 OS thread，執行 G 的實體 |
| P (Processor) | 負責調度 G 到 M 的中介者（最多 GOMAXPROCS 個）|

- G → P → M 的排程順序
- 每個 P 有自己的 Local Run Queue
- Global Run Queue：用於跨 P 之間的工作分派

---

## ⏱️ 調度原則與行為

- 非 preemptive（非強制搶佔式，舊版 runtime）
- Runtime 定時插入 safe-point（現在支援非搶佔式中斷）
- 長時間佔用 CPU 的 G 若不釋放 control，會影響其他 G 的排程
- `runtime.Gosched()` 可手動讓出 CPU 控制權

---

## ⚠️ 常見問題與陷阱

| 問題類型 | 範例 / 行為 |
|----------|--------------|
| goroutine 泄漏 | 忘記關閉 channel 或無限阻塞導致 goroutine 無法退出 |
| 資料競爭 | 多個 goroutine 同時讀寫共享變數，需加鎖處理 |
| 鎖死 / 死鎖 | 多個 goroutine 卡在 channel 或 Mutex 無法繼續 |
| 調度失控 | 無限制啟動 goroutine，造成爆炸性成長與 OOM |

---

## 💡 最佳實務建議

- ✅ 搭配 `context.Context` 控制 goroutine 壽命
- ✅ 每個 goroutine 都應該能「退出」
- ✅ 控制 goroutine 數量（例如 worker pool）
- ✅ 注意 channel 的讀寫對應關係
- ✅ 用 `sync.WaitGroup` 控制 goroutine 的結束時機

---

## 🧪 補充：如何觀察 goroutine 狀態

```go
runtime.NumGoroutine()   // 取得當前 goroutine 數量
pprof.Lookup("goroutine").WriteTo(os.Stdout, 1) // 堆疊輸出
```