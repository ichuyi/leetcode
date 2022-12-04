package main

func minSubArrayLen(target int, nums []int) int {
	min := 0
	windows := make([]int, 0)
	sum := 0
	for _, v := range nums {
		windows = append(windows, v)
		sum += v
		for sum >= target {
			l := len(windows)
			if l < min || min == 0 {
				min = l
			}
			first := windows[0]
			sum -= first
			windows = windows[1:l]
		}
	}
	return min
}
