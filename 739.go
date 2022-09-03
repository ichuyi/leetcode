package main

func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	offset := -1
	result := make([]int, len(temperatures))
	for i, t := range temperatures {
		for offset >= 0 && t > temperatures[stack[offset]] {
			result[stack[offset]] = i - stack[offset]
			stack = stack[:offset]
			offset--
		}
		stack = append(stack, i)
		offset++
		continue
	}
	return result
}
