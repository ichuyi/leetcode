package main

// MinStack TODO
type MinStack struct {
	data         []int
	dataOffset   int
	helper       []int
	helperOffset int
}

// Constructor TODO
// func Constructor() MinStack {
// 	return MinStack{
// 		data:         make([]int, 0),
// 		helper:       make([]int, 0),
// 		dataOffset:   -1,
// 		helperOffset: -1,
// 	}
// }

// Push TODO
func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	this.dataOffset++
	if this.helperOffset < 0 || val <= this.helper[this.helperOffset] {
		this.helper = append(this.helper, val)
		this.helperOffset++
	}
}

// Pop TODO
func (this *MinStack) Pop() {
	if this.dataOffset < 0 {
		return
	}
	t := this.data[this.dataOffset]
	this.data = this.data[:this.dataOffset]
	this.dataOffset--
	if this.helper[this.helperOffset] == t {
		this.helper = this.helper[:this.helperOffset]
		this.helperOffset--
	}
}

// Top TODO
func (this *MinStack) Top() int {
	if this.dataOffset < 0 {
		return 0
	}
	return this.data[this.dataOffset]
}

// GetMin TODO
func (this *MinStack) GetMin() int {
	if this.helperOffset < 0 {
		return 0
	}
	return this.helper[this.helperOffset]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
