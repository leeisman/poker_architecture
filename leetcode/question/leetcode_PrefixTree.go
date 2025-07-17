// 208. Implement Trie (Prefix Tree)
// 可以快速搜尋前綴，也適合敏感詞之類

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			isEnd:    false,
		},
	}
}

func (t *Trie) Insert(word string) {
	curr := t.root
	for _, ch := range word {
		if _, ok := curr.children[ch]; !ok {
			curr.children[ch] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		curr = curr.children[ch]
	}
	curr.isEnd = true
}

func (t *Trie) Search(word string) bool {
	curr := t.root
	for _, ch := range word {
		if _, ok := curr.children[ch]; !ok {
			return false
		}
		curr = curr.children[ch]
	}
	return curr.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t.root
	for _, ch := range prefix {
		if _, ok := curr.children[ch]; !ok {
			return false
		}
		curr = curr.children[ch]
	}
	return true
}