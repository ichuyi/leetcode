package main

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] > points[j][0]
	})
	r := len(points)
	min := points[0][0]
	ans := 1
	for i := 1; i < r; i++ {
		v := points[i]
		if v[1] < min {
			min = v[0]
			ans++
		}
	}
	return ans
}
