package main

import "fmt"

type LRUCache struct {
	capacity   int
	size       int
	cache      map[int]*ListNode
	head, tail *ListNode
}

type ListNode struct {
	key   int
	value int
	prev  *ListNode
	next  *ListNode
}

func addNode(lru *LRUCache, node *ListNode) {
	//add a new node at head
	node.prev = lru.head
	node.next = lru.head.next

	lru.head.next.prev = node
	lru.head.next = node
}

func popTail(lru *LRUCache) *ListNode {
	res := lru.tail.prev
	removeNode(res)
	return res
}

func removeNode(node *ListNode) {
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev
}

func moveToHead(lru *LRUCache, node *ListNode) {
	removeNode(node)
	addNode(lru, node)
}

func Constructor(capacity int) LRUCache {
	h := &ListNode{}
	t := &ListNode{}
	h.next = t
	t.prev = h
	return LRUCache{
		capacity: capacity,
		size:     0,
		cache:    map[int]*ListNode{},
		head:     h,
		tail:     t,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	moveToHead(this, node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if !ok {
		newNode := &ListNode{key: key, value: value}
		this.cache[key] = newNode
		addNode(this, newNode)
		this.size++
		if this.size > this.capacity {
			tail := popTail(this)
			delete(this.cache, tail.key)
			this.size--
		}
	} else {
		node.value = value
		moveToHead(this, node)
	}
}

func main() {

	lru := Constructor(4)
	fmt.Println(lru.Get(5))
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	lru.Put(5, 5)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(2))
	fmt.Println(lru.Get(5))
}
