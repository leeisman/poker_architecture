# Golang GMP 調度模型全解 (Goroutine, M, P)

## 目錄
- 什麼是 GMP 模型
- G/M/P 是什麼？
- GMP 的調度流程
- GOMAXPROCS 與 P 關係
- GC 與 goroutine 關聯性
- pprof/火焰圖 分析行為
- 優化實施措施

---

## 什麼是 GMP 模型
GMP 是 Go runtime 的調度模型，用來管理 goroutine（軟體執行單元）與 OS thread 之間的執行與調度關係。

---

## G/M/P 是什麼？

| 符號 | 內容 |
|------|------|
| G (Goroutine) | Go 軟體線程，由 runtime 管理 |
| M (Machine) | OS 執行緒，實際執行 G 的載體 |
| P (Processor) | 虛擬處理器，持有執行環境與 G 任務佇列 |

---

## GMP 調度流程

```mermaid
sequenceDiagram
    autonumber
    participant App as 使用程式
    participant G as Goroutine (G)
    participant RunQ as 本地 P 任務佇列
    participant P as Processor (P)
    participant M as OS Thread (M)
    participant Runtime as Go Runtime

    App->>Runtime: go func()
    Runtime->>G: 建立 G
    G->>RunQ: 加入 local runq
    M->>RunQ: 從 runq 取 G
    RunQ->>M: 回傳 G
    M->>G: 執行 goroutine
    G-->>M: 結束 / 停任
    M->>Runtime: 回收 / 重新調度
```
---
## OS Thread 調度（傳統做法，如 Java）
```mermaid
sequenceDiagram
    autonumber
    participant Task1
    participant Task2
    participant Task3
    participant OS_Thread1 as OS Thread 1
    participant OS_Thread2 as OS Thread 2

    Task1->>OS_Thread1: 建立 thread 處理 Task1
    Task2->>OS_Thread2: 建立 thread 處理 Task2
    Task3->>OS: 等待 thread 空閒
    Note over OS_Thread1,OS_Thread2: 切換 thread（context switch）交由 OS 決定
```
---
##  Go GMP 模型調度
```mermaid
sequenceDiagram
    autonumber
    participant G1 as Goroutine 1
    participant G2 as Goroutine 2
    participant G3 as Goroutine 3
    participant P1 as Processor (P)
    participant M1 as OS Thread (M)

    G1->>P1: 放入 P1 的 run queue
    G2->>P1: 放入 P1 的 run queue
    G3->>P1: 放入 P1 的 run queue
    Note over P1: Go 調度器決定執行順序
    P1->>M1: 把 G1 給 M1 執行
    M1->>G1: 執行 G1（無需進入 OS）
    P1->>M1: 再給 G2
    M1->>G2: 執行 G2
```
---
## 三色標記
```mermaid
sequenceDiagram
    autonumber
    participant G as Goroutine
    participant Stack as G Stack
    participant Pool as sync.Pool
    participant GC as Go GC Runtime
    participant Buf as []byte (Buffer)

    Note over G,Pool: 開始使用 sync.Pool

    G->>Pool: Get() 取得 Buf
    Pool-->>G: 傳回 Buf（舊 pool 中物件）

    G->>Stack: 將 Buf 存到 goroutine stack 中
    Note over GC: GC 啟動！標記階段開始

    GC->>Stack: 從 GC Root 開始掃描（包含 goroutine stack）
    GC->>Buf: 發現被引用 → 標記為灰色
    GC->>Buf: 掃描結束 → 標記為黑色（不可回收）

    G->>Pool: 使用完 Buf → Put() 回 pool

    Note over Pool: pool.next = current → 切換 pool 區段

    Note over GC: 下一輪 GC 開始時，清除上輪沒被取出的物件
    GC->>Pool: 檢查舊區段的 pool（pool.victims）
    alt Buf 沒人再用（stack 無引用）
        GC->>Buf: 標記為白色 → 被釋放
    else Buf 在 stack 或 global 被引用
        GC->>Buf: 標記為灰 → 黑，保留不釋放
    end
```
```mermaid
graph TD
    A[CPU cacheline]
    B[Thread stack 上變數引用物件]
    C[GC Root 指向物件]
    D[GC 標記為灰 → 黑]
    E[Heap 中沒被引用的區塊]
    F[GC 結束後掃描 → 白色物件]
    G[OS Page 分配 Bitmap 清除]

    B --> C --> D
    E --> F --> G
    A --> B
```