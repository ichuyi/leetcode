package main

import "math"

func myAtoi(s string) int {
	p := 0
	firstBit := 0
	numbers := make([]int, 0)
	for _, v := range s {
		if v == '-' && p == 0 {
			p = -1
			firstBit++
			continue
		}
		if v == '+' && p == 0 {
			p = 1
			firstBit++
			continue
		}
		if v == ' ' && firstBit == 0 {
			continue
		}
		if v < '0' || v > '9' {
			break
		}
		if p == 0 {
			p = 1
		}
		firstBit++
		numbers = append(numbers, int(v)-int('0'))
	}
	var result int64 = 0
	for _, v := range numbers {
		result = result*10 + int64(v)
		if result*int64(p) > int64(math.MaxInt32) {
			return math.MaxInt32
		}
		if result*int64(p) < int64(math.MinInt32) {
			return math.MinInt32
		}
	}
	return int(result) * p
}
