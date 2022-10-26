package main

func removeKdigits(num string, k int) string {
	stack := make([]rune, 0)
	offset := -1
	for _, v := range num {
		for offset > -1 && k > 0 && stack[offset] > v {
			stack = stack[:offset]
			offset--
			k--
		}
		stack = append(stack, v)
		offset++
	}
	stack = stack[0 : offset-k+1]
	start := 0
	for _, v := range stack {
		if v != '0' {
			break
		}
		start++
	}
	if start == len(stack) {
		return "0"
	} else {
		return string(stack[start:])
	}
}
