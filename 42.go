package main

func trap(height []int) int {
	stack := make([]int, 0)
	offset := -1
	sum := 0
	for i, v := range height {
		for offset >= 1 && v >= height[stack[offset]] {
			left, top := stack[offset-1], stack[offset]
			sum += (min(v, height[left]) - height[top]) * (i - left - 1)
			stack = stack[:offset]
			offset--
		}
		for offset >= 0 && v >= height[stack[offset]] {
			stack = stack[:offset]
			offset--
		}
		stack = append(stack, i)
		offset++
	}
	return sum
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
