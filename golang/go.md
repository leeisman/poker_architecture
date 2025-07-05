# Golang 面試常見技術題總整理

---

## 🧠 一、Golang 語言核心（必考）⭐️⭐️⭐️⭐️⭐️

| 主題              | 面試常見問題                                     | 建議複習重點                              |
|-------------------|--------------------------------------------------|-------------------------------------------|
| Goroutine         | goroutine 是怎麼調度的？與 OS thread 差異？     | GMP 模型、M:N 調度、runtime.Gosched()     |
| Channel           | buffered vs unbuffered 差異？如何防止 deadlock？| select 多路複用、close 行為               |
| defer             | 執行順序？與 return 調用順序有關嗎？             | 後進先出、defer 先於 return 評估參數     |
| slice / map       | append 是否會改變底層？map 是否安全？            | slice capacity、map 非併發安全性          |
| 指標與值接收者    | 接收者該用值還是指標？                           | 逃逸分析、修改需求、性能差異              |
| interface         | interface 如何運作？type assertion 怎麼寫？     | itab、type switch、空 interface 原理      |

---

## 🔍 二、進階語言特性（中高考）⭐️⭐️⭐️⭐️

| 主題              | 面試常見問題                                  | 建議複習重點                               |
|-------------------|-----------------------------------------------|--------------------------------------------|
| 逃逸分析          | 什麼情況會從 stack 逃逸到 heap？              | `go build -gcflags=-m` 使用方法            |
| sync 套件         | sync.Once、Pool、Mutex 差在哪？               | buffer 重用、鎖、只執行一次                |
| context           | 如何取消 context？怎麼設計鏈式 context？      | context.WithCancel / WithTimeout          |
| panic / recover   | 怎麼捕捉 panic？應用在哪？                    | defer + recover 使用時機與陷阱            |
| unsafe / reflect  | 為什麼用？有哪些使用風險？                    | memory layout、struct tag、性能損耗        |

---

## 🧰 三、標準庫與工具（常考）⭐️⭐️⭐️⭐️

| 主題              | 面試常見問題                                 | 建議複習重點                               |
|-------------------|----------------------------------------------|--------------------------------------------|
| net/http          | 怎麼寫 middleware？如何設定 context？        | ServeMux、HandlerFunc、context 傳遞        |
| json/xml 編解碼   | `omitempty`、`inline` 是什麼意思？            | struct tag 用法、字段忽略、深層嵌套        |
| testing 套件      | 如何寫單元測試？怎麼 mock？                  | table-driven、gomock、httptest             |
| go mod            | go.mod 怎麼管理版本？replace 的作用？        | private module、replace 使用時機           |
| 編譯與優化        | 怎麼交叉編譯？如何減少 GC 壓力？             | pprof 分析、`GOOS/GOARCH` 設定              |

---

## ⚙️ 四、系統設計與架構（資深必考）⭐️⭐️⭐️⭐️⭐️

| 主題              | 面試常見問題                                 | 建議複習重點                                |
|-------------------|----------------------------------------------|---------------------------------------------|
| 微服務架構        | 如何做服務註冊與發現？retry / timeout 如何設計？| ZK/Consul、service mesh vs SDK 呼叫         |
| 高併發設計        | 如何解決 race condition？是否用 atomic？     | sync vs atomic、channel 處理並發邏輯       |
| 無狀態設計        | 為什麼無狀態重要？怎麼做水平擴充？           | Redis 外部狀態管理、冪等處理設計            |
| GC 特性與優化     | Golang 的 GC 怎麼運作？怎麼避免 GC spike？   | 三色標記、heap 增長、高頻 GC 分析           |
| 訊息隊列 / Kafka  | Kafka 為什麼可靠？如何設計 retry？           | offset 管理、partition 設計、at-least-once |

---

## 📦 五、部署與實戰經驗（高分加分題）⭐️⭐️⭐️

| 主題              | 面試常見問題                               | 建議複習重點                                 |
|-------------------|--------------------------------------------|----------------------------------------------|
| CI/CD 流程        | 你們如何部署？如何滾動更新不中斷？         | git-flow、feature flag、blue-green          |
| 容器化與 K8s      | 有無用 Docker？健康檢查怎麼設計？          | Dockerfile、readinessProbe/livenessProbe    |
| 效能監控          | 如何找出效能瓶頸？                          | pprof 使用、Prometheus + Grafana 整合        |