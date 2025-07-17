# 📘 Staff Engineer LeetCode 練習清單

這份清單著重於系統設計、資料結構深度、效能敏感性，適合具備中高資歷的工程師強化核心實力與面試準備。

---

## ✅ 快取與資料淘汰設計

| 題號 | 題目 | 技術核心 | 練習重點 |
|------|------|-----------|-----------|
| 146 | [LRU Cache](https://leetcode.com/problems/lru-cache/) | Doubly Linked List + HashMap | 模擬快取淘汰策略、設計可擴充結構 |
| 460 | [LFU Cache](https://leetcode.com/problems/lfu-cache/) | 頻率統計 + 雙層結構 | 記憶體敏感系統的資料淘汰策略 |

---

## ✅ 即時資料處理 / 統計

| 題號 | 題目 | 技術核心 | 練習重點 |
|------|------|-----------|-----------|
| 295 | [Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/) | MinHeap + MaxHeap | 建構支援即時統計的資料結構 |
| 215 | [Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/description/)| MinHeap + MaxHeap | 建構支援 priority queue|
---

## ✅ 字典查詢與搜尋優化

| 題號 | 題目 | 技術核心 | 練習重點 |
|------|------|-----------|-----------|
| 208 | [Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/) | Tree + Map 結構 | 快速字串查詢、Autocomplete 模型 |
| 981 | [Time-Based Key-Value Store](https://leetcode.com/problems/time-based-key-value-store/) | 時序資料查詢 + Binary Search | 快取版本紀錄、TTL 設計模擬 |

---

## ✅ 字串 / 資料抽取技巧

| 題號 | 題目 | 技術核心 | 練習重點 |
|------|------|-----------|-----------|
| 76 | [Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/) | Sliding Window + 頻率表 | 訊息擷取 / 資料掃描效率優化 |

---

## ✅ 隨機資料操作

| 題號 | 題目 | 技術核心 | 練習重點 |
|------|------|-----------|-----------|
| 380 | [Insert Delete GetRandom O(1)](https://leetcode.com/problems/insert-delete-getrandom-o1/) | HashMap + Array | 設計 O(1) 存取與隨機抽樣 |

---

## ✅ 併發與同步控制

| 題號 | 題目 | 技術核心 | 練習重點 |
|------|------|-----------|-----------|
| 1114 | [Print in Order](https://leetcode.com/problems/print-in-order/) | Goroutine 模擬 | 基礎 Concurrency 控制、Golang 實作練習 |

---

## ✅ 拓撲與依賴管理

| 題號 | 題目 | 技術核心 | 練習重點 |
|------|------|-----------|-----------|
| 210 | [Course Schedule II](https://leetcode.com/problems/course-schedule-ii/) | 拓撲排序 + DAG | 任務依賴管理、模擬部署順序 |

---

## ✅ 跨系統整併 / 合併排序

| 題號 | 題目 | 技術核心 | 練習重點 |
|------|------|-----------|-----------|
| 23 | [Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/) | MinHeap + Linked List | 多資料來源整併 / 合併排序設計 |

---

## 🧠 自訂挑戰（進階 Staff 等級）

- **實作 Consistent Hashing**（模擬分散式系統 routing）
- **設計一個 Rate Limiter**（token bucket 或 sliding log）
- **設計 Goroutine-safe LRU**（考慮 concurrency 與鎖）

---