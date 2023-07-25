package main

func longestPalindrome(str string) string {
	ans := ""
	for i := 0; i < len(str); i++ {
		l, r := find(str, i, i)
		if len(ans) < r-l {
			ans = string(str[l:r])
		}
		if i+1 < len(str) {
			l, r = find(str, i, i+1)
			if len(ans) < r-l {
				ans = string(str[l:r])
			}
		}

	}
	return ans
}

// find 左闭右开
func find(str string, l, r int) (int, int) {
	for l >= 0 && r < len(str) {
		if str[l] != str[r] {
			return l + 1, r
		}
		l--
		r++
	}
	return l + 1, r
}
