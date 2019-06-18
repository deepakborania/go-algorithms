package main

import "fmt"

type Trie struct {
	root *Node
}

type Node struct {
	valid bool
	nodes []*Node
}

func constructNode() *Node {
	return &Node{
		valid: false,
		nodes: make([]*Node, 26),
	}
}

func Constructor() Trie {
	return Trie{
		root: constructNode(),
	}
}

func (this *Trie) Insert(word string) {
	curr := this.root
	for i := 0; i < len(word); i++ {
		if i == len(word)-1 {
			curr.valid = true
		}
		idx := word[i] - 'a'
		if curr.nodes[idx] == nil {
			curr.nodes[idx] = constructNode()
		}
		curr = curr.nodes[idx]
	}
}

func (this *Trie) Search(word string) bool {
	curr := this.root
	for i := 0; i < len(word); i++ {
		if i == len(word)-1 && curr.valid {
			return true
		}
		idx := word[i] - 'a'
		curr = curr.nodes[idx]
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for i := 0; i < len(prefix); i++ {
		idx := prefix[i] - 'a'
		if i == len(prefix)-1 && curr.nodes[idx] != nil {
			return true
		}
		curr = curr.nodes[idx]
	}
	return false
}

func main() {
	t := Constructor()
	t.Insert("abc")
	t.Insert("abcra")
	t.Insert("abcar")
	t.Insert("def")
	// fmt.Println(t.Search("abc"))
	// fmt.Println(t.Search("ab"))
	// fmt.Println(t.Search("abcar"))
	// fmt.Println(t.Search("abca"))
	// fmt.Println(t.Search("abcra"))
	fmt.Println("")
	fmt.Println(t.StartsWith("abc"))
	fmt.Println(t.StartsWith("abs"))
}
