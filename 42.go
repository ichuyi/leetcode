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

func trap2(height []int) int {
	leftMax, rightMax := make([]int, len(height)), make([]int, len(height))
	leftMax[0], rightMax[len(height)-1] = height[0], height[len(height)-1]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}
	ans := 0
	for i, v := range height {
		ans += min(leftMax[i], rightMax[i]) - v
	}
	return ans
}
func trap3(height []int) int {
	left, right := 1, len(height)-2
	leftMax, rightMax := height[left-1], height[right+1]
	ans := 0
	for left < right {
		if leftMax < rightMax {
			leftMax = max(leftMax, height[left])
			ans += leftMax - height[left]
			left++
		} else {
			rightMax = max(rightMax, height[right])
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}
