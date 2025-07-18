# TiDB Staff Engineer 學習筆記總覽

本筆記涵蓋 TiDB 架構、分散式原理、交易處理、效能優化與實戰應用，適合 Staff Engineer 面試與大型系統導入前的深度學習。

---

## 📄 文件目錄與說明

### 01_TiDB_Architecture_Overview.md  
TiDB 整體架構、元件介紹（TiDB Server、TiKV、PD、TiFlash），與 MySQL、Vitess 的對比優勢。

### 02_TiDB_Region_and_Sharding.md  
Region 是什麼？如何基於 Primary Key sharding？Region 分裂、遷移與 Leader 分配邏輯。

### 03_TiDB_Transaction_Model.md  
Percolator 模型、2PC 實作、樂觀/悲觀鎖、分散式事務的一致性保證。

### 04_TiDB_Cluster_Deployment.md  
使用 TiUP 部署、滾動升級、多副本與可用區設計、節點擴容與故障修復流程。

### 05_TiDB_SQL_and_Index_Design.md  
主鍵與唯一鍵策略、避免 Hot Region、Partition Table 設計原則與使用場景。

### 06_TiDB_Query_Optimization.md  
EXPLAIN ANALYZE 使用、查詢下推（predicate/aggregation）、TiFlash 使用時機與優化方式。

### 07_TiDB_Performance_and_Monitoring.md  
常見慢查詢排查、Prometheus + Grafana 指標觀察、Region/Store 分佈熱點分析。

### 08_TiDB_Reliability_and_Backup.md  
多副本一致性、Raft 共識、BR/Dumpling 備份還原策略與跨雲容災備援。

### 09_TiDB_Use_Cases.md  
金流系統設計（錢包 + 交易紀錄）、遊戲事件記錄、高併發寫入的 schema 設計案例。

### 10_TiDB_Comparison_with_MySQL.md  
與 MySQL 8.0 在效能、擴展性、交易模型的實際差異與場景選擇比較。

---

## 📘 建議學習順序  
TiDB_Architecture_Overview → Region_and_Sharding → Transaction_Model → SQL_and_Index_Design → Optimization → Reliability → Use_Cases

---