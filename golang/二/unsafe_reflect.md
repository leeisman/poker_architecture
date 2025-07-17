
## Go 面試筆記：unsafe / reflect（Staff Engineer 適用）

---

### ✅ 面試常見問題

| 問題                     | 答案 |
|--------------------------|------|
| 為什麼要使用 unsafe？   | 操作記憶體位址、提升效能、達成低階優化（如 memory layout 調整） |
| 為什麼要使用 reflect？  | 動態取得型別資訊、操作 struct tag、做泛型資料處理（但效能較差） |
| 使用有哪些風險？         | unsafe 會破壞型別安全；reflect 效能低、結構複雜，易造成 runtime bug |

---

### ✅ unsafe 常見用途

- 改變 struct memory layout 排序（壓縮結構體大小）
- 轉型：`uintptr` <-> `unsafe.Pointer`
- 提升 cache locality 避免 padding
- 模擬 union 型別
- 零拷貝 string <-> []byte 轉換
- 快速欄位存取（跳過反射）

```go
// float64 -> uint64 bit 操作
var f float64 = 3.14
p := (*uint64)(unsafe.Pointer(&f))
fmt.Println(*p)

// 零拷貝 string -> []byte
func StringToBytes(s string) []byte {
    strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
    return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
        Data: strHeader.Data,
        Len:  strHeader.Len,
        Cap:  strHeader.Len,
    }))
}
```

---

### ✅ reflect 常見用途

- 解析 JSON struct tag（搭配 encoding/json）
- 泛型處理：如驗證器、ORM、自動轉換器
- 搭配 `TypeOf()`、`ValueOf()` 做通用方法

```go
t := reflect.TypeOf(myStruct)
for i := 0; i < t.NumField(); i++ {
    field := t.Field(i)
    fmt.Println("Field Name:", field.Name, "Tag:", field.Tag)
}
```

---

### ✅ 反射替代場景：高效能 JSON 解析

```go
// 使用 unsafe 快速設定欄位
type Player struct {
    Name string
}

func setName(ptr unsafe.Pointer, offset uintptr, val string) {
    fieldPtr := unsafe.Pointer(uintptr(ptr) + offset)
    *(*string)(fieldPtr) = val
}
```

---

### ✅ 建議與風險

| 項目 | 建議 |
|------|------|
| ✅ 適用場景 | low-level 優化 / 編碼空間節省（unsafe）動態邏輯解耦（reflect） |
| ❌ 不建議 | 商業邏輯濫用 unsafe、過度依賴反射製造複雜性 |
| 陷阱 | 不小心破壞對齊與類型安全導致記憶體錯誤 |

---

### ✅ 小結

- `unsafe` 是效能優化利器，但需嚴格控制使用位置與測試
- 常用於零拷貝轉換、欄位快速寫入、記憶體壓縮
- `reflect` 提供彈性但成本高，適合工具層做泛型抽象
- 面試可結合場景，如記憶體壓縮設計、ORM 解析、protobuf 編碼器等

---

