
## Golang Interface 完整筆記

### ✅ Interface 是什麼？
- 是一組方法的集合：任何實作了 interface 所有方法的型別，都被視為「實作了這個 interface」
- 在 Go 中屬於「隱式實作」（Implicit implementation），**不用明確宣告 implements**

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

---

### ✅ Interface 的應用場景
1. **抽象設計 / 解耦依賴**
   - 如：`io.Reader`, `http.Handler`
2. **單元測試 mock**
   - 用 interface 替換實體，方便 mock 依賴行為
3. **策略模式 / 插拔式實作**
   - 可根據不同環境傳入不同實作

---

### ✅ 實作規則（隱式）

```go
type Animal interface {
    Speak() string
}

type Dog struct{}
func (d Dog) Speak() string {
    return "Woof"
}

var a Animal = Dog{} // ✅ OK：Dog 有實作 Speak()
```

---

### ⚠️ 值與指標實作差異
- 如果 interface 定義的方法是透過 **指標接收者** 實作，則只有 `*T` 能滿足 interface

```go
type Doer interface {
    Do()
}

type Worker struct{}
func (w *Worker) Do() {} // 只有 *Worker 實作了 Doer

var d Doer = &Worker{} // ✅ OK
var d2 Doer = Worker{}  // ❌ 編譯錯誤
```

---

### ✅ 空介面（interface{}）
- 可接受任意型別（等同於 Java 的 Object）
- 實務應用：JSON 解析、泛型參數、map[string]interface{}

```go
func PrintAny(val interface{}) {
    fmt.Println(val)
}
```

---

### 🧠 Interface 內部原理簡述
| 名稱 | 說明 |
|------|------|
| itab | method table，記錄型別對應的 method 實作 |
| data | 指向實際儲存的值（值或指標） |

當你寫 `var r io.Reader = file`，其實是創建了一個 `interface{itab, data}` 結構，透過 itab 執行對應方法。

---

### 📌 常見面試題與回答

**1. interface 為什麼是隱式實作？優點是什麼？**  
答：這讓 interface 的使用更靈活，解耦性高，使用者不需侵入式修改型別。只要方法對得上，就能自動被認定為實作 interface，提升模組化與可測試性。

**2. 為什麼空介面能接受所有型別？**  
答：因為在 Go 中，所有型別都隱式實作了空 interface（interface{} 沒有方法），所以任何東西都可以當成 interface{} 傳遞。

**3. interface 如何與值/指標接收者互動？有何限制？**  
答：如果你用指標接收者實作 interface 方法，只有指標型別能被賦值給 interface。若用值接收者實作，則值與指標都可以被賦值。

**4. interface 是否可以為 nil？要注意什麼陷阱？**  
答：可以，但要小心 `interface != nil` 並不代表裡面沒包 nil。`interface{itab, data}` 的 itab 不為 nil 時，整個 interface != nil，即使 data 是 nil。

---

### ❗ interface == nil 陷阱
```go
var r io.Reader = (*os.File)(nil)
fmt.Println(r == nil) // false！
```

- 因為 `r` 包含了 itab（非 nil）和 data（是 nil）
- 所以 `interface != nil` 成立，即使裡面包的是 nil
- 正確做法：`if r == nil || reflect.ValueOf(r).IsNil() { ... }`

---

### ✅ 總結：面試重點心法
| 主題 | 重點說明 |
|------|----------|
| 隱式實作 | 不需 implements，靠方法簽名比對 |
| 指標 vs 值 | 若 interface 需求的方法是指標接收者，必須用 *T 來實作 |
| 空介面 | 用於泛型、JSON、多型包裝 |
| nil 陷阱 | interface 本身不為 nil 但內部 data 可能為 nil，需小心判斷 |

---

掌握 interface 的本質與限制，是 Go 高階工程師非常重要的一環。建議實際寫 mock、抽象層設計、泛型包裝等應用來加深理解。
