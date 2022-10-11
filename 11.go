package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ma := 0
	for left < right {
		h := min(height[left], height[right])
		ma = max(ma, h*(right-left))
		if height[left] < height[right] {
			left--
		} else {
			right--
		}
	}
	return ma
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
