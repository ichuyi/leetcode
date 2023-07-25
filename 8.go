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
func myAtoi2(s string) int {
	nums := make([]int, 0)
	ws := false
	p := 1
	pf := false
	for _, v := range s {
		if v == ' ' {
			if ws {
				break
			} else {
				continue
			}
		}
		ws = true
		if v == '-' {
			if pf {
				break
			}
			p = -1
			pf = true
			continue
		}
		if v == '+' {
			if pf {
				break
			}
			p = 1
			pf = true
			continue
		}
		pf = true
		if v < '0' || v > '9' {
			break
		}
		nums = append(nums, int(v-'0')*p)
	}
	ans := 0
	for _, v := range nums {
		ans = ans*10 + v
		if ans > math.MaxInt32 {
			ans = math.MaxInt32
		}
		if ans < math.MinInt32 {
			ans = math.MinInt32
		}
	}
	return ans
}
