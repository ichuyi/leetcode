package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	gi, si := 0, 0
	for gi < len(g) && si < len(s) {
		if g[gi] <= s[si] {
			gi++
			si++
			continue
		}
		si++
	}
	return gi
}
