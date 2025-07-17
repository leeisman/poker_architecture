// 295 Find Median From Data Stream
// 設計 MinHeap + MaxHeap 
// 透過Addnum 將數值分類到正確的heap並且排序
// 在最終從minHeap + maxHeap取得值就是中位數

// 計heap

// 確認實作interface
var _ heap.Interface = (*IntHeap)(nil)

type IntHeap strcut{
    data []Int
    isMinHeap bool
}

func NewIntHeap(isMinHeap bool) *IntHeap{
    h:= &IntHeap{
        data: make([]int),
        isMinHeap: isMinHeap
    }
    heap.Init(h)
    return h
}

func (h *IntHeap) Len() int{ len(h.data)}
func (h *IntHeap) Less(i,j int) bool{
    if h.isMinHeap{
        return h.data[i] < h.data[j]
    }
    return h.data[i] > h.data[j]
}
func (h *IntHeap) Swap(i,j int){
    h.data[i], h.data[j] = h.data[j], h.data[i]
}

func(h *IntHeap) Push(x Interface){
    h.data = append(h.data, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[: n-1]
	return x
}

// 自定義函數
func (h *IntHeap) Top() (int, bool){
    if h.Len() == 0{
        return 0, false
    }
    return h.data[0], true
}

func (h *IntHeap) PishInt(x int){
    heap.Push(h, x)
}

func (h *IntHeap) PopInt() (int, bool){
    if h.Len() == 0{
        return 0 , false
    }
    return head.Pop(h).(int), true
}

// -----------------------
// MedianFinder
// -----------------------

type MedianFinder struct {
    left *IntHeap  // max-heap
    right *IntHeap // min-heap
}

func NewMedianFinder() *MedianFinder{
    return &MedianFinder{
        left: NewIntHeap(false),
        right: NewIntHeap(true),
    }
}

func (mf *MedianFinder) AddNum(num int) {
	// 先加入 maxHeap（左邊）
	if mf.left.Len() == 0 || num <= mf.left.data[0] {
		mf.left.PushInt(num)
	} else {
		mf.right.PushInt(num)
	}

	// 平衡兩邊大小
	if mf.left.Len() > mf.right.Len()+1 {
		val, _ := mf.left.PopInt()
		mf.right.PushInt(val)
	} else if mf.right.Len() > mf.left.Len() {
		val, _ := mf.right.PopInt()
		mf.left.PushInt(val)
	}
}

func (m *MedianFinder) FindMedian() float64{
    if m.left.Len() > m.right.Len(){
        val, _:= m.left.Top()
        return float64(val)
    }
    leftTop, _ := m.left.Top()
    rightTop, _ := m.right.Top()
    return float64(leftTop+rightTop) / 2.0
}

func Constructor() MedianFinder {
	return MedianFinder{
		left:  NewIntHeap(false),
		right: NewIntHeap(true),
	}
}