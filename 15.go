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
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for i > 0 && i < len(nums) && nums[i] == nums[i-1] {
			i++
		}
		if i == len(nums) {
			break
		}
		j := i + 1
		target := -nums[i]
		for k := len(nums) - 1; j < k; {
			sum := nums[j] + nums[k]
			if sum > target {
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
				if j == k {
					break
				}
			} else if sum < target {
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				if j == k {
					break
				}
			} else {
				// res = append(res, []int{nums[i], nums[j], nums[k]})
				res = append(res, []int{i, j, k})
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return res
}
