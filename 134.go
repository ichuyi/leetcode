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

func canCompleteCircuit2(gas []int, cost []int) int {
	start := 0
	l := len(gas)
	for start < l {
		var i int
		sum := gas[start]
		for i = 0; i < l; i++ {
			index := (start + i) % l
			sum = sum - cost[index]
			if sum < 0 {
				break
			}
			sum += gas[(index+1)%l]
		}
		if i == l {
			return start
		}
		start = start + i + 1
	}
	return -1
}
