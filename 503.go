package main

func nextGreaterElements(nums []int) []int {
	l := len(nums)
	result := make([]int, l)
	for i := range result {
		result[i] = -1
	}
	stack := make([]int, 0)
	offset := -1
	index := 0
	left := l
	for left > 0 {
		for offset >= 0 && nums[index%l] > nums[stack[offset]] {
			result[stack[offset]] = nums[index%l]
			left--
			stack = stack[:offset]
			offset--
		}
		if offset > -1 && stack[offset] == index%l {
			stack = stack[:offset]
			left--
			offset--
			index--
		}
		if index < l {
			stack = append(stack, index)
			offset++
		}
		index = index + 1
	}
	return result
}
