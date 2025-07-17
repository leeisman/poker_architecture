// 146 LRU Cache
// 主要了解 doubly linked list,
// 透過結構 LRUCache 實作get put函數
// get時候將節點移除並且放置最前面
// 超過capacity的時候將tail的前一個移除

head <> A <> B <> C <> tail
type listNode struct {
    key, val int
	prev, next *listNode
}

type LRUCache struct{
	cache map[int]*listNode
	head, tail *listNode
	capacity int
}

func Construcstor(capacity int) *LRUCache{
	head := &listNode
	tail := &listNode
	head.next = tail
	tail.prev = head
	return &LRUCache{
		head:head,
		tail:tail,
		capacity:capacity,
		cache: make(map[int]*listNode)
	}
}

func (this *LRUCache)Get(key int){
	if node,ok := this.cache[key];ok{
		this.moveToFront(node)
		return node.val
	}

	return -1
}

func (this *LRUCache)Put(key,value int){
	if node,ok := this.cache[key];ok{
		node.val = value
		this.moveToFront(node)
		return
	}

	node := &listNode{key:key, val:value}
	this.cache[key] = node
	this.addToFront(node)

	if len(this.cache) > this.capacity{
		lru := this.removeTail()
		delete(this.cache, lru.key)
	}
}

// head <> A
func(this *LRUCache)addToFront(node *listNode){
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func(this *LRUCache)moveToFront(node *listNode){
	this.removeNode(node)
	this.addToFront(node)
}

// prev <> node <> next
func(this *LRUCache)removeNode(node *listNode){
	node.prev.next = node.next
	node.next.prev = node.prev
}

func(this *LRUCache)removeTail() *listNode{
	tail := this.tail.prev
	this.removeNode(tail)
	return tail
}