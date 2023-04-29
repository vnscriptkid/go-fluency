package abuse

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		child, ok := node.children[ch]
		if !ok {
			child = NewTrieNode()
			node.children[ch] = child
		}
		node = child
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		child, ok := node.children[ch]
		if !ok {
			return false
		}
		node = child
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		child, ok := node.children[ch]
		if !ok {
			return false
		}
		node = child
	}
	return true
}

type fastAbuseMasker struct {
	trie *Trie
}

func NewFastAbuseMasker() IAbuseMasking {
	m := &fastAbuseMasker{
		trie: NewTrie(),
	}

	for _, word := range abusiveWordsList {
		m.trie.Insert(word)
	}

	return m
}

func (m *fastAbuseMasker) Mask(message string) string {
	maskedMessage := []rune(message)

	for i := 0; i < len(maskedMessage); i++ {
		node := m.trie.root
		for j, ch := range maskedMessage[i:] {
			child, ok := node.children[ch]
			if !ok {
				break
			}

			if child.isEnd {
				for k := i; k <= i+j; k++ {
					maskedMessage[k] = '*'
				}
			}
			node = child
		}
	}

	return string(maskedMessage)
}
