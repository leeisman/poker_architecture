
# 🧩 Golang Slice 與 Map 重點整理

---

## 🍰 Slice 是什麼？

- Slice 是基於 array 的動態大小序列，是 Go 中最常用的資料結構之一。
- 包含三個欄位：**pointer、length、capacity**
- 宣告方式：
  ```go
  var s []int
  s := []int{1, 2, 3}
  s := make([]int, 3, 5) // len = 3, cap = 5
  ```

---

## 🧪 Slice 操作與注意事項

- 透過 `append()` 增加元素，可能導致底層陣列重建
- 使用 `copy()` 可複製 slice
- Slice 傳遞的是 **引用**，修改會影響原本資料
- 擷取語法：`s[1:3]`，左含右不含
- 擷取不會複製資料，仍是共用底層陣列

```go
a := []int{1, 2, 3, 4}
b := a[1:3] // b = [2 3]
b[0] = 9
fmt.Println(a) // [1 9 3 4] → 原 a 被改變
```

---

## ⚠️ Slice 常見陷阱

| 問題               | 原因或說明                               |
|--------------------|------------------------------------------|
| 容量不足時 append  | 會建立新陣列，原 slice 內容可能不同步     |
| 傳參導致誤修改     | slice 傳遞是引用語意，會影響原資料         |
| 擷取後持續引用     | 造成潛在記憶體洩漏（原 array 無法 GC 回收） |

---

## 🧠 Map 是什麼？

- Map 是 key-value 對應的集合，內建雜湊表結構
- 宣告方式：
  ```go
  m := map[string]int{"a": 1, "b": 2}
  m := make(map[string]int)
  ```

---

## 🛠️ Map 操作

| 操作              | 說明                     |
|-------------------|--------------------------|
| `m[k]`            | 取得 key 對應的值        |
| `m[k] = v`        | 設定或新增 key-value     |
| `delete(m, k)`    | 刪除 key                 |
| `v, ok := m[k]`   | 檢查 key 是否存在         |

```go
val, ok := m["foo"]
if ok {
    fmt.Println("found:", val)
} else {
    fmt.Println("not found")
}
```

---

## ⚠️ Map 陷阱與限制

| 陷阱或限制            | 說明                                     |
|-----------------------|------------------------------------------|
| 非執行緒安全           | 多 goroutine 同時讀寫需加鎖              |
| key 必須是可比對型別   | 如：string、int、struct（不可包含 slice） |
| map 是引用型別         | 傳遞或賦值都是共享相同 map               |
| 不保證順序             | `range` map 的順序是隨機的               |

---

## 🧠 面試常見題

- slice 與 array 差在哪？
- slice 傳遞參數會不會改到原本的？
- append 後 slice 為什麼變了？
- map 為什麼不是 thread-safe？
- map 的底層結構是什麼？為什麼 key 不能是 slice？

---

> 📌 小提醒：slice 和 map 都是「引用型別」，傳遞時會共用底層資料，使用時要特別注意修改行為與記憶體使用。
