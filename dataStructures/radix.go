package main

import (
	"fmt"
)

type RadixTrie struct {
	root *RadixTrieNode
}

type RadixTrieNode struct {
	value       rune
	children    []*RadixTrieNode
	isEndOfWord bool
}

func NewRadixTrie() *RadixTrie {
	return &RadixTrie{root: &RadixTrieNode{value: ' '}}
}

func (t *RadixTrie) Insert(word string) {
	if word == "" {
		return
	}
	current := t.root
	for _, char := range word {
		child := t.findChild(current, char)
		if child == nil {
			child = &RadixTrieNode{value: char}
			current.children = append(current.children, child)
		}
		current = child
	}
	current.isEndOfWord = true
}

func (t *RadixTrie) Search(word string) bool {
	current := t.root
	for _, char := range word {
		child := t.findChild(current, char)
		if child == nil {
			return false
		}
		current = child
	}
	return current.isEndOfWord
}

func (t *RadixTrie) findChild(node *RadixTrieNode, char rune) *RadixTrieNode {
	for _, child := range node.children {
		if child.value == char {
			return child
		}
	}
	return nil
}

func main() {
	trie := NewRadixTrie()
	trie.Insert("apple")
	trie.Insert("banana")
	trie.Insert("orange")
	fmt.Println(trie.Search("apple"))  // Output: true
	fmt.Println(trie.Search("banana")) // Output: true
	fmt.Println(trie.Search("orange")) // Output: true
	fmt.Println(trie.Search("pear"))   // Output: false
}
