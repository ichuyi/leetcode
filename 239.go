package main

import (
	"container/heap"
	"sort"
)

type hp struct {
	sort.IntSlice
	a []int
}

// Less TODO
func (h hp) Less(i, j int) bool { return h.a[h.IntSlice[i]] > h.a[h.IntSlice[j]] }

// Push TODO
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }

// Pop TODO
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func maxSlidingWindow(nums []int, k int) []int {
	q := make([]int, 0, k)
	for i := 0; i < k-1; i++ {
		q = append(q, i)
	}
	h := hp{IntSlice: q, a: nums}
	heap.Init(&h)
	result := make([]int, 0)
	for j := k - 1; j < len(nums); j++ {
		heap.Push(&h, j)
		for h.IntSlice[0] < (j - k + 1) {
			heap.Pop(&h)
		}
		m := h.IntSlice[0]
		result = append(result, nums[m])
	}
	return result
}
func maxSlidingWindow2(nums []int, k int) []int {
	q := make([]int, 0)
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	for i := 0; i < k-1; i++ {
		push(i)
	}
	result := make([]int, 0)
	for i := k - 1; i < len(nums); i++ {
		push(i)
		for q[0] < i-k+1 {
			q = q[1:]
		}
		result = append(result, nums[q[0]])
	}
	return result
}
