package main

import (
    "fmt"
    "net/http"
)

func Add(x, y int) (res int) {
	return x + y
}

// Routes interface implementation using a Radix
// trie data structure
type TrieRoutes Trie

func NewTrieRoutes() TrieRoutes {
    trie := NewTrie()
    return TrieRoutes{trie.root}
}

type Trie struct {
	root *node
}

type node struct {
	children map[rune]*node
	isEnd    bool
    handler  map[httpMethod]*http.HandlerFunc
}

func NewTrie() *Trie {
	return &Trie{
		root: &node{
			children: make(map[rune]*node),
			isEnd:    false,
            handler:  nil,
		},
	}
}

func (t *Trie) Insert(word string, h http.HandlerFunc) {
	currentNode := t.root
	for _, c := range word {
		if _, ok := currentNode.children[c]; !ok {
			currentNode.children[c] = &node{
				children: make(map[rune]*node),
				isEnd:    false,
                handler:  nil,
			}
		}
		currentNode = currentNode.children[c]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(word string) bool {
	currentNode := t.root
	for _, c := range word {
		if _, ok := currentNode.children[c]; !ok {
			return false
		}
		currentNode = currentNode.children[c]
	}
	return currentNode.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	currentNode := t.root
	for _, c := range prefix {
		if _, ok := currentNode.children[c]; !ok {
			return false
		}
		currentNode = currentNode.children[c]
	}
	return true
}

func (t *Trie) Print() {
    t.printHelper(t.root, []rune{})
}

func (t *Trie) printHelper(node *node, word []rune) {
    if node.isEnd {
        fmt.Println(string(word))
    }
    for ch, child := range node.children {
        t.printHelper(child, append(word, ch))
    }
}

func main () {
    trie := NewTrie()
    handler := func(w http.ResponseWriter, r *http.Request) {}
    trie.Insert("apple", handler)
    trie.Insert("application", handler)

    trie.Print()

    fmt.Println(trie.Search("app"))
    fmt.Println(trie.StartsWith("app"))
}

