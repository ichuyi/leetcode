package main

func firstMissingPositive(nums []int) int {
	for index := 0; index < len(nums); {
		if nums[index] <= 0 || nums[index] > len(nums) || nums[index] == index+1 || nums[index] == nums[nums[index]-1] {
			index++
		} else {
			nums[index], nums[nums[index]-1] = nums[nums[index]-1], nums[index]
		}
	}
	for i, v := range nums {
		if v != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
