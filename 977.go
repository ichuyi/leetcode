package main

func sortedSquares(nums []int) []int {
	i, j := 0, len(nums)-1
	result := make([]int, len(nums))
	index := len(nums) - 1
	for i <= j {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			result[index] = nums[i] * nums[i]
			i++
		} else {
			result[index] = nums[j] * nums[j]
			j--
		}
		index--
	}
	return result
}
