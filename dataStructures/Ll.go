package main

import (
    "fmt"
)
type TrieNode struct {
    children map[rune]*TrieNode
    isEnd    bool
}

type Trie struct {
    root *TrieNode
}

func NewTrie() *Trie {
    return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
    node := t.root
    for _, ch := range word {
        if _, ok := node.children[ch]; !ok {
            node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
        }
        node = node.children[ch]
    }
    node.isEnd = true
}

func (t *Trie) Search(word string) bool {
    node := t.root
    for _, ch := range word {
        if _, ok := node.children[ch]; !ok {
            return false
        }
        node = node.children[ch]
    }
    return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
    node := t.root
    for _, ch := range prefix {
        if _, ok := node.children[ch]; !ok {
            return false
        }
        node = node.children[ch]
    }
    return true
}
func (t *Trie) Print() {
    t.printNode(t.root, "")
}

func (t *Trie) printNode(node *TrieNode, prefix string) {
    if node == nil {
        return
    }

    if node.isEnd {
        fmt.Println(prefix)
    }

    for ch, child := range node.children {
        t.printNode(child, prefix+string(ch))
    }
}


func main() {
    trie := NewTrie()
    trie.Insert("apple")
    trie.Insert("application")


    fmt.Println("\nPrint:")

    trie.Print()

    fmt.Println()


    fmt.Printf("Search app -> %v\n", trie.Search("app"))   // Output: true
    fmt.Printf("StartsWith app -> %v\n", trie.StartsWith("app"))   // Output: true
}
