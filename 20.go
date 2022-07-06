package main

func isValid(s string) bool {
	m := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	l := make([]rune, len(s))
	cur := -1
	for _, v := range s {
		if _, ok := m[v]; ok {
			cur++
			l[cur] = v
			continue
		}
		if cur < 0 || m[l[cur]] != v {
			return false
		}
		cur--
	}
	return cur == -1
}

// func main() {
// 	fmt.Println(isValid("([])"))
// 	fmt.Println(isValid("()[]"))
// 	fmt.Println(isValid("(]"))
// }
