package main

import (
	"sort"
)

func merge56(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		v := intervals[i]
		if v[0] > right {
			result = append(result, []int{left, right})
			left, right = v[0], v[1]
		} else {
			if v[1] > right {
				right = v[1]
			}
		}
	}
	result = append(result, []int{left, right})
	return result
}

func merge562(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := make([][]int, 0)
	reset := true
	start, end := 0, 0
	for i := 0; i < len(intervals); {
		cur := intervals[i]
		if reset {
			start, end = cur[0], cur[1]
			reset = false
			i++
			continue
		}
		if end < cur[0] {
			ans = append(ans, []int{start, end})
			reset = true
			continue
		}
		if cur[1] > end {
			end = cur[1]
		}
		i++
	}
	if !reset {
		ans = append(ans, []int{start, end})
	}
	return ans
}
