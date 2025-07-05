
# 🪝 Golang defer 重點整理

---

## 🧠 defer 是什麼？

- `defer` 是 Golang 的延遲執行語法
- 表示在 **函數返回之前執行的語句**
- 常用於：資源釋放、解鎖、紀錄、錯誤處理

```go
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
// 輸出順序：hello → world
```

---

## 🧾 執行時機與特性

| 特性               | 說明 |
|--------------------|------|
| 執行時機           | 在**當前函數 return 之前**執行 |
| 後進先出（LIFO）   | 多個 defer 會逆序執行         |
| 參數立即求值       | defer 語句宣告當下就計算參數，不等到真正執行時 |

```go
func test() {
    x := 1
    defer fmt.Println(x) // 這裡 x 已經是 1
    x = 10
}
```

---

## 🔁 多個 defer 的執行順序（LIFO）

```go
func main() {
    defer fmt.Println("A")
    defer fmt.Println("B")
    defer fmt.Println("C")
}
// 輸出：C → B → A
```

---

## ⚠️ defer 搭配 return 的細節

```go
func foo() int {
    x := 5
    defer func() {
        x += 1
    }()
    return x // 傳回值為 5，不是 6
}
```

### 📌 原因：
- `return x` 是**值已經被複製**後，才執行 defer。
- 除非 return 是命名回傳值，才可能被 defer 修改。

```go
func bar() (x int) {
    defer func() {
        x += 1
    }()
    return 5 // 回傳值變成 6
}
```

---

## 🔐 常見用途範例

### ✅ 解鎖資源

```go
mu.Lock()
defer mu.Unlock()
```

### ✅ 關閉檔案或連線

```go
f, _ := os.Open("file.txt")
defer f.Close()
```

### ✅ 回收 panic、錯誤處理

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from panic:", r)
    }
}()
```

---

## ⚠️ 常見陷阱與注意事項

| 問題類型 | 說明與範例 |
|----------|------------|
| ✅ defer 立即求值 | 傳入的參數當下就被求值，不是等到執行 defer 時才算 |
| ❌ defer 在迴圈內 | 容易造成資源堆疊，建議使用閉包函數避免 |
| ⚠️ 執行效率 | defer 有一定效能開銷，頻繁呼叫可能影響效能 |

```go
for i := 0; i < 1000; i++ {
    defer fmt.Println(i) // 不建議這樣用
}
```

---

## 🧠 面試常見題

- defer 的執行順序是？
- defer 的參數是何時求值？
- defer 能否修改 return 的值？
- 在哪裡使用 defer 最適合？
- defer 有效能影響嗎？
- 如何利用 defer 做錯誤處理？

---

## 🔍 補充：與 panic/recover 搭配使用

```go
func main() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("panic 被捕捉:", err)
        }
    }()

    panic("爆炸啦")
}
```

> 📌 defer + recover 是 Golang 中實作 try-catch-like 的方式
