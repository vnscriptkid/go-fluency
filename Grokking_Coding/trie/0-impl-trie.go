package Trie

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

func (t *TrieNode) InsertHere(str string, i int) *TrieNode {
	if i == len(str) {
		return t
	}

	// cur char
	c := rune(str[i])

	// check children, if cur char exists in children
	var node *TrieNode

	if node = t.children[c]; node == nil {
		node = NewTrieNode()

		t.children[c] = node
	}

	if i == len(str)-1 {
		node.isWord = true
	}

	// insert next char at node
	node.InsertHere(str, i+1)

	return t
}

// NewTrieNode() is a contructor function
func NewTrieNode() *TrieNode {
	node := new(TrieNode)
	node.isWord = false
	node.children = make(map[rune]*TrieNode)

	for char := 'a'; char <= 'z'; char++ {
		node.children[char] = nil
	}

	return node
}

type Trie struct {
	// Write your code here
	Root *TrieNode
}

// constructor initializes Trie type object
func NewTrie() *Trie {
	// Your code will replace this placeholder return statement
	return &Trie{
		Root: NewTrieNode(),
	}
}

// Insert function insert string in trie
func (t *Trie) Insert(word string) {
	t.Root.InsertHere(word, 0)
}

// Search function search for a string
func (t *Trie) Search(word string) bool {
	return t.Root.SearchDFS(word, 0, false)
}

func (n *TrieNode) SearchDFS(word string, i int, checkPrefix bool) bool {
	if i == len(word) {
		if checkPrefix {
			return true
		}

		return n.isWord
	}

	curChar := rune(word[i])

	var node *TrieNode

	if node = n.children[curChar]; node == nil {
		return false
	}

	return node.SearchDFS(word, i+1, checkPrefix)
}

// SearchPrefix function search for a prefix
func (t *Trie) SearchPrefix(prefix string) bool {
	return t.Root.SearchDFS(prefix, 0, true)
}
