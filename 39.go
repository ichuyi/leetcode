package main

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	return doCombinationSum(candidates, target, nil, nil)
}
func doCombinationSum(candidates []int, target int, tmp []int, res [][]int) [][]int {
	if target == 0 {
		r := make([]int, len(tmp))
		copy(r, tmp)
		res = append(res, r)
		return res
	}
	if len(candidates) == 0 || target <= 0 {
		return res
	}
	for i, v := range candidates {
		if v <= target {
			res = doCombinationSum(candidates[i:], target-v, append(tmp, v), res)
		} else {
			break
		}
	}
	return res
}
