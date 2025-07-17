// 460 LFU Cache

// LFUCache
// ├── capacity = 3
// ├── size = 2
// ├── minFreq = 1
// ├── nodes = map[
// │     10 => &Node{key:10, val:100, freq:1},
// │     20 => &Node{key:20, val:200, freq:2}
// │   ]
// └── freqs = map[
//       1 => DoublyLinkedList [10]
//       2 => DoublyLinkedList [20]
//     ]

type Node struct {
	key, val, freq int
	prev, next     *Node
}

type DoublyLinkedList struct {
	head, tail *Node
}

func newDoublyLinkedList() *DoublyLinkedList {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &DoublyLinkedList{head, tail}
}

func (dll *DoublyLinkedList) isEmpty() bool {
	return dll.head.next == dll.tail
}

func (dll *DoublyLinkedList) addToFront(node *Node) {
	node.next = dll.head.next
	node.prev = dll.head
	dll.head.next.prev = node
	dll.head.next = node
}

func (dll *DoublyLinkedList) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (dll *DoublyLinkedList) removeLast() *Node {
	if dll.isEmpty() {
		return nil
	}
	node := dll.tail.prev
	dll.removeNode(node)
	return node
}
type LFUCache struct {
	capacity int
	size     int
	minFreq  int
	nodes    map[int]*Node
	freqs    map[int]*DoublyLinkedList
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		nodes:    make(map[int]*Node),
		freqs:    make(map[int]*DoublyLinkedList),
	}
}

func (c *LFUCache) Get(key int) int {
	if node, ok := c.nodes[key]; ok {
		c.increaseFreq(node)
		return node.val
	}
	return -1
}

func (c *LFUCache) Put(key int, value int) {
	if c.capacity == 0 {
		return
	}
	if node, ok := c.nodes[key]; ok {
		node.val = value
		c.increaseFreq(node)
	} else {
		if c.size == c.capacity {
			toRemove := c.freqs[c.minFreq].removeLast()
			delete(c.nodes, toRemove.key)
			c.size--
		}
		newNode := &Node{key: key, val: value, freq: 1}
		c.nodes[key] = newNode
		if c.freqs[1] == nil {
			c.freqs[1] = newDoublyLinkedList()
		}
		c.freqs[1].addToFront(newNode)
		c.minFreq = 1
		c.size++
	}
}

func (c *LFUCache) increaseFreq(node *Node) {
	oldFreq := node.freq
	c.freqs[oldFreq].removeNode(node)
	if c.freqs[oldFreq].isEmpty() {
		delete(c.freqs, oldFreq)
		if c.minFreq == oldFreq {
			c.minFreq++
		}
	}
	node.freq++
	if c.freqs[node.freq] == nil {
		c.freqs[node.freq] = newDoublyLinkedList()
	}
	c.freqs[node.freq].addToFront(node)
}