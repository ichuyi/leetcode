package main

func search(nums []int, target int) int {
	return doSearch(nums, 0, len(nums)-1, target)
}
func doSearch(nums []int, low, high, target int) int {
	if low > high {
		return -1
	}
	if nums[low] == target {
		return low
	}
	if nums[high] == target {
		return high
	}
	mid := (low + high) / 2
	if target == nums[mid] {
		return mid
	}
	if nums[low] < nums[mid] {
		// 左边有序
		if target >= nums[low] && target < nums[mid] {
			return doSearch(nums, low, mid-1, target)
		}
		return doSearch(nums, mid+1, high, target)
	} else {
		// 右边有序
		if target > nums[mid] && target <= nums[high] {
			return doSearch(nums, mid+1, high, target)
		}
		return doSearch(nums, low, mid-1, target)
	}
}
