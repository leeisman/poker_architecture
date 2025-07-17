
## Golang Escape Analysis（逃逸分析）筆記

### ✅ 什麼是 Escape Analysis？
Escape Analysis 是 Go 編譯器用來判斷變數應該分配在 stack 還是 heap 的靜態分析技術。

---

### 🧠 Stack vs Heap 的差異
| 項目 | Stack | Heap |
|------|-------|------|
| 分配速度 | 快（函式內） | 慢（需 GC 管理） |
| 回收方式 | 自動（出 stack 即回收） | 需垃圾回收（GC） |
| 存活時間 | 短（隨著函式結束） | 長（跨函式、goroutine） |

---

### ✅ 為什麼會發生 Escape？
Go 編譯器在以下情況判斷變數「可能活得太久」，就會將其分配到 heap：

1. 被函式 return 回去（外部仍要用）
2. 被儲存在 interface、slice、map 等 reference type 中
3. 被閉包（closure）捕捉
4. 傳遞給需要指標參數的函式，且其生命週期不明

---

### ✅ 如何查看是否逃逸？
使用 `go build -gcflags="-m"` 可輸出編譯器分析：

```bash
go run -gcflags="-m" main.go
```

範例輸出：
```
main.go:10:6: moved to heap: user
main.go:15:10: &user escapes to heap
```

---

### ✅ 範例說明
```go
type User struct {
    Name string
}

func NewUser(name string) *User {
    u := User{Name: name}
    return &u // u 被返回 → 逃逸到 heap
}
```

---

### ✅ 如何避免不必要的逃逸
1. 避免不必要的指標傳遞
2. 使用 struct 傳值（小結構）
3. 將變數限制在區域作用域內
4. 減少 interface 包裝

---

### ⚠️ 實務陷阱
| 行為 | 說明 |
|------|------|
| `fmt.Println(&val)` | 傳指標進 fmt 系列會逃逸（因為 fmt 接收 interface{}） |
| map[string]*T | value 是 pointer，也會強制逃逸 |
| slice of pointer | 一樣會逃逸，因為 slice 可能超出作用域 |

---

### ✅ 面試常見問題與回答
**Q1: Go 的變數什麼時候會逃逸到 heap？**  
A1: 當編譯器認為變數離開當前作用域仍然可能被使用時，例如 return 指標、傳給 closure、interface 包裝等情況。

**Q2: 為什麼逃逸會影響效能？**  
A2: heap 分配速度慢且會增加 GC 負擔，尤其在大量短生命週期物件時。

**Q3: 怎麼 debug 我的變數有沒有逃逸？**  
A3: 使用 `go build -gcflags=-m`，觀察變數是否有 "moved to heap" 的訊息。

---

### ✅ 總結
- Stack 分配速度快、GC 低負擔
- Heap 分配可跨函式使用，但效能較差
- 逃逸是由編譯器自動判斷，開發者可透過 `-gcflags=-m` 協助分析
- 儘可能讓變數侷限在函式內並避免 interface 包裝，可提升效能

掌握逃逸分析，是寫出高效 Golang 程式碼的重要基礎
