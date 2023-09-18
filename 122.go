package main

func maxProfit(prices []int) int {
	m := 0
	buy := -1
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			if buy == -1 {
				buy = prices[i]
			}
		} else if prices[i] > prices[i+1] && buy > -1 {
			m += prices[i] - buy
			buy = -1
		}
	}
	if buy > -1 {
		m += prices[len(prices)-1] - buy
	}
	return m
}
func maxProfit1(prices []int) int {
	sum := 0
	buy := -1
	for i := 0; i < len(prices)-1; i++ {
		cur := prices[i]
		next := prices[i+1]
		if cur > next {
			if buy != -1 {
				sum += cur - buy
				buy = -1
			}
			continue
		}
		if buy == -1 {
			buy = cur
		}
	}
	if buy != -1 {
		sum += prices[len(prices)-1] - buy
	}
	return sum
}
