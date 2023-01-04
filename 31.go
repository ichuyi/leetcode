package main

func nextPermutation(nums []int) {
	j := len(nums) - 2
	for j >= 0 && nums[j] >= nums[j+1] {
		j--
	}
	if j >= 0 {
		k := len(nums) - 1
		for k > j && nums[k] <= nums[j] {
			k--
		}
		nums[j], nums[k] = nums[k], nums[j]
	}
	for i, k := j+1, len(nums)-1; i < k; {
		nums[i], nums[k] = nums[k], nums[i]
		i++
		k--
	}
}
