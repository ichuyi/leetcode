package main

// MyQueue TODO
type MyQueue struct {
	first        []int
	firstOffset  int
	second       []int
	secondOffset int
}

// Constructor TODO
func Constructor() MyQueue {
	return MyQueue{
		firstOffset:  -1,
		secondOffset: -1,
	}
}

// Push TODO
func (this *MyQueue) Push(x int) {
	this.first = append(this.first, x)
	this.firstOffset++
}

// Pop TODO
func (this *MyQueue) Pop() int {
	n := this.Peek()
	this.second = this.second[0:this.secondOffset]
	this.secondOffset--
	return n
}

// Peek TODO
func (this *MyQueue) Peek() int {
	if this.secondOffset < 0 {
		for this.firstOffset >= 0 {
			this.second = append(this.second, this.first[this.firstOffset])
			this.first = this.first[0:this.firstOffset]
			this.firstOffset--
			this.secondOffset++
		}
	}
	if this.secondOffset >= 0 {
		n := this.second[this.secondOffset]
		return n
	}
	return 0
}

// Empty TODO
func (this *MyQueue) Empty() bool {
	return this.firstOffset < 0 && this.secondOffset < 0
}
