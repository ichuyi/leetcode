package main

func canCompleteCircuit(gas []int, cost []int) int {
	oil := make([]int, 0, len(gas))
	for i := range gas {
		oil = append(oil, gas[i]-cost[i])
	}
	l := len(oil)
	start := 0
	for start < l {
		s := oil[start]
		if s < 0 {
			start++
			continue
		}
		for index := 1; index < l; index++ {
			i := (index + start) % l
			s += oil[i]
			if s < 0 {
				start = (index + start + 1)
				break
			}
		}
		if s >= 0 {
			return start
		}
	}
	return -1
}
