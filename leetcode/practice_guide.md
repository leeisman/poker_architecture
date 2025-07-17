# 資料結構總覽表（LeetCode 練習指南）

| 資料結構    | 定義與特性                | 常見操作               | 時間複雜度（平均）       | LeetCode 常見題型            |
|:------------|:--------------------------|:-----------------------|:-------------------------|:-----------------------------|
| Array       | 固定大小、連續記憶體      | 隨機存取快，插入刪除慢 | 查找 O(1)、插刪 O(n)     | Two Pointers, Sliding Window |
| Linked List | 節點指向下一個，不連續    | 插入刪除快，查找慢     | 查找 O(n)、插刪 O(1)     | 反轉鏈表、環形偵測、合併排序 |
| String      | 字元陣列，常為 immutable  | 子字串、比對、轉換     | 查找/比較 O(n)           | KMP、Z-algorithm、字串壓縮   |
| Stack       | LIFO，後進先出            | push / pop             | O(1)                     | 括號配對、單調棧、後序表達式 |
| Queue       | FIFO，先進先出            | enqueue / dequeue      | O(1)                     | BFS、滑動視窗、最短路徑      |
| Heap (PQ)   | 完全二元樹，支援找極值    | 插入 / 取最值          | O(log n)                 | Top K 頻率、合併區間         |
| HashMap     | key-value 映射表          | 插入、查詢、刪除       | O(1)（平均）             | Two Sum、LRU、字元統計       |
| Matrix      | 2D 陣列                   | 掃描、鄰近判斷         | O(m×n)                   | 圖像處理、DFS/BFS 模擬       |
| Tree        | 節點結構 + 遞迴關係       | 遍歷、插入、平衡       | 查找 O(log n)、遍歷 O(n) | DFS、BFS、BST、DP on Tree    |
| Graph       | 頂點 + 邊（鄰接表或矩陣） | 遍歷、最短路徑         | O(V+E)                   | DFS、BFS、拓撲排序、Dijkstra |