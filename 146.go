package main

// DoubleNode TODO
type DoubleNode struct {
	Key, Val   int
	Next, Last *DoubleNode
}

// LRUCache TODO
type LRUCache struct {
	root, tail *DoubleNode
	m          map[int]*DoubleNode
	capacity   int
}

// ConstructorLRU TODO
func ConstructorLRU(capacity int) LRUCache {
	root := &DoubleNode{}
	tail := &DoubleNode{}
	root.Next = tail
	tail.Last = root
	return LRUCache{
		root:     root,
		tail:     tail,
		m:        make(map[int]*DoubleNode, capacity),
		capacity: capacity,
	}
}

// Get TODO
func (this *LRUCache) Get(key int) int {
	v, ok := this.m[key]
	if !ok {
		return -1
	}
	this.remove(v)
	this.moveToHead(v)
	return v.Val
}

// Put TODO
func (this *LRUCache) Put(key int, value int) {
	v, ok := this.m[key]
	if ok {
		v.Val = value
		this.remove(v)
	} else {
		v = &DoubleNode{Key: key, Val: value}
		this.m[key] = v
	}
	this.moveToHead(v)
	if len(this.m) <= this.capacity {
		return
	}
	remove := this.tail.Last
	delete(this.m, remove.Key)
	this.remove(remove)
}
func (this *LRUCache) moveToHead(n *DoubleNode) {
	next := this.root.Next
	this.root.Next = n
	n.Next = next
	next.Last = n
	n.Last = this.root
}
func (this *LRUCache) remove(n *DoubleNode) {
	last := n.Last
	next := n.Next
	last.Next = next
	next.Last = last
	n.Next = nil
	n.Last = nil
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
