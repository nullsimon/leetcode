package lrucache

type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

type LRUCache struct {
	Capacity int
	Size     int
	Cache    map[int]*Node

	Head *Node
	Tail *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	lru := LRUCache{
		Capacity: capacity,
		Size:     0,
		Cache:    make(map[int]*Node),
		Head:     head,
		Tail:     tail,
	}
	return lru
}

func (lru *LRUCache) Get(key int) int {
	if lru.Size == 0 {
		return -1
	}
	node, ok := lru.Cache[key]
	if !ok {
		return -1
	}
	//move node to front
	lru.MoveBack(node)
	return node.Value
}

func (lru *LRUCache) Put(key int, value int) {

	n, ok := lru.Cache[key]
	if ok {
		n.Value = value
		lru.MoveBack(n)
		return
	}

	lru.Size++
	node := &Node{
		Key:   key,
		Value: value,
	}
	lru.Cache[key] = node
	tail := lru.Tail
	if tail == nil {
		lru.Head = node
		lru.Tail = node
		return
	}
	if tail != nil {
		tail.Next = node
		node.Prev = tail
		node.Next = nil
		lru.Tail = node
	}
	if lru.Size > lru.Capacity {
		lru.Delete(lru.Head)
	}
}

func (lru *LRUCache) Delete(node *Node) {
	if lru.Size == 0 {
		return
	}
	node, ok := lru.Cache[node.Key]
	if !ok {
		return
	}
	if lru.Size == 1 {
		lru.Head = nil
		lru.Tail = nil
	} else {
		next := lru.Head.Next
		next.Prev = nil
		lru.Head = next
	}
	delete(lru.Cache, node.Key)
	lru.Size--

}

func (lru *LRUCache) MoveBack(node *Node) {
	if lru.Size == 1 {
		return
	}

	//tail already problems
	if node.Next == nil {
		return
	}

	//head
	if node.Prev == nil {
		nextNode := node.Next
		nextNode.Prev = nil
		if nextNode.Next == nil {
			nextNode.Next = node
		}
		lru.Head = nextNode

		node.Next = nil
		node.Prev = lru.Tail
		lru.Tail = node
		return
	}

	//not head and tail, middle
	prevNode := node.Prev
	nextNode := node.Next

	prevNode.Next = nextNode

	nextNode.Prev = prevNode
	if nextNode.Next == nil {
		nextNode.Next = node
	}

	node.Prev = lru.Tail
	node.Next = nil
	lru.Tail = node

}
