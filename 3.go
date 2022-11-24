package main

func lengthOfLongestSubstring(s string) int {
	window := make([]rune, 0)
	set := make(map[rune]struct{})
	max := 0
	for _, v := range s {
		_, ok := set[v]
		for ok {
			c := window[0]
			delete(set, c)
			window = window[1:]
			_, ok = set[v]
		}
		window = append(window, v)
		set[v] = struct{}{}
		if max < len(window) {
			max = len(window)
		}
	}
	return max
}
