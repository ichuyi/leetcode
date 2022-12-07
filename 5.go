package main

func longestPalindrome(str string) string {
	start, end := 0, 0
	for i := 0; i < len(str); i++ {
		s, e := f(str, i, i)
		if e-s > end-start {
			start, end = s, e
		}
		if i < len(str)-1 && str[i] == str[i+1] {
			s, e = f(str, i, i+1)
			if e-s > end-start {
				start, end = s, e
			}
		}
	}
	return str[start : end+1]
}

func f(s string, i, j int) (int, int) {
	if i <= 0 || j >= len(s)-1 {
		return i, j
	}
	if s[i-1] == s[j+1] {
		return f(s, i-1, j+1)
	} else {
		return i, j
	}
}
