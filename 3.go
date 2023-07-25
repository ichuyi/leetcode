package main

func lengthOfLongestSubstring(s string) int {
	window := make(map[rune]struct{})
	start := 0
	ans := 0
	for i, v := range s {
		for _, ok := window[v]; ok; _, ok = window[v] {
			vv := s[start]
			delete(window, rune(vv))
			start++
		}
		if _, ok := window[v]; !ok {
			window[v] = struct{}{}
			ans = max(ans, i-start+1)
			continue
		}
	}
	return ans
}
