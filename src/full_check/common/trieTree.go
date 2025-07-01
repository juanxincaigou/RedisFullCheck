package common

type TrieNode struct {
	children map[byte]*TrieNode
	isEnd    bool // 是否为精确匹配
	isPrefix bool // 是否为前缀匹配
	isSuffix bool // 是否为后缀匹配
	isAll    bool // 是否为全匹配（空串或 "*"）
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode),
		isEnd:    false,
		isPrefix: false,
		isSuffix: false,
		isAll:    false,
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: newTrieNode()}
}

func reverse(word []byte) []byte {
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
	return word
}

func (trie *Trie) Insert(word []byte) {
	node := trie.root

	// 空串或 "*" 视为全匹配
	if len(word) == 0 || (len(word) == 1 && word[0] == '*') {
		node.isAll = true
		return
	}

	// 前缀匹配
	if word[len(word)-1] == '*' {
		for _, char := range word[:len(word)-1] {
			if _, ok := node.children[char]; !ok {
				node.children[char] = newTrieNode()
			}
			node = node.children[char]
		}
		node.isPrefix = true
		return
	}

	// 后缀匹配
	if word[0] == '*' {
		reversedWord := reverse(word[1:])
		for _, char := range reversedWord {
			if _, ok := node.children[char]; !ok {
				node.children[char] = newTrieNode()
			}
			node = node.children[char]
		}
		node.isSuffix = true
		return
	}

	// 精确匹配
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			node.children[char] = newTrieNode()
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func (trie *Trie) Search(word []byte) bool {
	node := trie.root

	if node.isAll {
		return true
	}

	// 精确匹配
	curr := node
	for _, char := range word {
		if _, ok := curr.children[char]; !ok {
			curr = nil
			break
		}
		curr = curr.children[char]
	}
	if curr != nil && curr.isEnd {
		return true
	}

	// 前缀匹配
	curr = node
	for _, char := range word {
		if curr.isPrefix {
			return true
		}
		if _, ok := curr.children[char]; !ok {
			curr = nil
			break
		}
		curr = curr.children[char]
	}
	if curr != nil && curr.isPrefix {
		return true
	}

	// 后缀匹配
	reversedWord := reverse(append([]byte(nil), word...))
	curr = node
	for _, char := range reversedWord {
		if curr.isSuffix {
			return true
		}
		if _, ok := curr.children[char]; !ok {
			return false
		}
		curr = curr.children[char]
	}
	if curr != nil && curr.isSuffix {
		return true
	}

	return false
}
