package main

// MyCircularDeque TODO
type MyCircularDeque struct {
	l []int
	c int
}

// Constructor TODO
// func Constructor(k int) MyCircularDeque {
// 	return MyCircularDeque{l: make([]int, 0, k), c: k}
// }

// InsertFront TODO
func (this *MyCircularDeque) InsertFront(value int) bool {
	if len(this.l) < this.c {
		n := make([]int, 0, len(this.l)+1)
		n = append(n, value)
		this.l = append(n, this.l...)
		return true
	}
	return false
}

// InsertLast TODO
func (this *MyCircularDeque) InsertLast(value int) bool {
	if len(this.l) < this.c {
		this.l = append(this.l, value)
		return true
	}
	return false
}

// DeleteFront TODO
func (this *MyCircularDeque) DeleteFront() bool {
	if len(this.l) > 0 {
		this.l = this.l[1:]
		return true
	}
	return false
}

// DeleteLast TODO
func (this *MyCircularDeque) DeleteLast() bool {
	if len(this.l) > 0 {
		this.l = this.l[:len(this.l)-1]
		return true
	}
	return false
}

// GetFront TODO
func (this *MyCircularDeque) GetFront() int {
	if len(this.l) > 0 {
		return this.l[0]
	}
	return -1
}

// GetRear TODO
func (this *MyCircularDeque) GetRear() int {
	if len(this.l) > 0 {
		return this.l[len(this.l)-1]
	}
	return -1
}

// IsEmpty TODO
func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.l) <= 0
}

// IsFull TODO
func (this *MyCircularDeque) IsFull() bool {
	return len(this.l) >= this.c
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
