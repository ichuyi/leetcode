package main

import "math"

func subSort(array []int) []int {
	if len(array) <= 1 {
		return []int{-1, -1}
	}
	left, right := -1, -1
	max, min := math.MinInt, math.MaxInt
	for i, v := range array {
		if v >= max {
			max = v
		} else {
			left = i
		}
	}
	for i := len(array) - 1; i >= 0; i-- {
		v := array[i]
		if v <= min {
			min = v
		} else {
			right = i
		}
	}
	return []int{right, left}
}
