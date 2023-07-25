package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	l := len(nums)
	r := math.MaxInt
	for i := 0; i < l; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := l - 1
		for j < k {
			s := nums[i] + nums[j] + nums[k]
			if s == target {
				return s
			}
			if abs(r-target) > abs(s-target) {
				r = s
			}
			if s > target {
				var k1 int
				for k1 = k - 1; j < k1 && nums[k] == nums[k1]; k1-- {
				}
				if j == k1 {
					break
				}
				k = k1
			} else {
				var j1 int
				for j1 = j + 1; j1 < k && nums[j1] == nums[j]; j1++ {
				}
				if j1 == k {
					break
				}
				j = j1
			}
		}
	}
	return r
}
func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
func threeSumClosest2(nums []int, target int) int {
	sort.Ints(nums)
	ans := math.MaxInt
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}
			if abs(sum-target) < abs(ans-target) {
				ans = sum
			}
			if sum > target {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < target {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}
			if j == k {
				break
			}
		}
	}
	return ans
}
