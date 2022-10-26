package main

func canJump(nums []int) bool {
	l := len(nums)
	right := nums[0]
	for current := 0; current <= right && current < l; current++ {
		if current+nums[current] > right {
			right = current + nums[current]
		}
		if right >= l {
			return true
		}
	}
	return false
}
