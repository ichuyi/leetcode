package main

import "math"

func maxProfit2(prices []int) int {
	r := 0
	buy := math.MaxInt32
	for _, v := range prices[:len(prices)-1] {
		r = max(r, v-buy)
		buy = min(buy, v)
	}
	return r
}
