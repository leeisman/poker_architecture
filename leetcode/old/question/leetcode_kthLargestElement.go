// 215 Kth Largest Element in an Array
// 設計 MinHeap

// 計heap

// 確認實作interface
var _ heap.Interface = (*MinHeap)(nil)

type MinHeap []int
type (h *MinHeap)Len{ return len(h)}
type (h *MinHeap)Less(i,j int) bool { return h[i]< h[j] }
type (h *MinHeap)Swap(i,j int) { h[i],h[j] = h[j],h[i]}
type (h *MinHeap)Push(x Interface){
    *h = append(*h, x.(int))
}
type (h *MinHeap)Pop()interface {
    old = *h
    n := len(old)
    val := old[n-1]
    *h = old[:n-1]
    return val
}

func findKthLargest(nums []int, k int) int{
    h := &MinHeap{}
    heap.Init(h)

    for _, num := range nums{
        heap.Push(h, num)
        if h.Len() > k{
            heap.Pop(h)
        }
    }

    return (*h)[0]
}