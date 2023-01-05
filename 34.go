package main

func searchRange(nums []int, target int) []int {
	return []int{getLeftRange(nums, 0, len(nums)-1, target), getRightRange(nums, 0, len(nums)-1, target)}
}

func getLeftRange(nums []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	midNum := nums[mid]
	if midNum < target {
		return getLeftRange(nums, mid+1, high, target)
	}
	left := getLeftRange(nums, low, mid-1, target)
	if midNum == target && left == -1 {
		return mid
	}
	return left
}
func getRightRange(nums []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	midNum := nums[mid]
	if midNum > target {
		return getRightRange(nums, low, mid-1, target)
	}
	right := getRightRange(nums, mid+1, high, target)
	if midNum == target && right == -1 {
		return mid
	}
	return right
}
