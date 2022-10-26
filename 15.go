package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	l := len(nums)
	result := make([][]int, 0)
	sort.Ints(nums)
	for first := 0; first < l; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		target := -nums[first]
		for second := first + 1; second < l-1; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			three := l - 1
			for second < three && nums[second]+nums[three] > target {
				three--
			}
			if second == three {
				break
			}
			if nums[second]+nums[three] == target {
				result = append(result, []int{nums[first], nums[second], nums[three]})
			}
		}
	}
	return result
}
