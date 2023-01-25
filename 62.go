package main

func uniquePaths(m int, n int) int {
	if m*n == 0 {
		return 0
	}
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	res[0][0] = 1
	doUniquePaths(m-1, n-1, m-1, n-1, res)
	return res[m-1][n-1]
}
func doUniquePaths(m, n, x, y int, res [][]int) {
	if res[x][y] > 0 {
		return
	}
	var r2, r1 int
	if x > 0 {
		if res[x-1][y] == 0 {
			doUniquePaths(m, n, x-1, y, res)
		}
		r1 = res[x-1][y]
	}
	if y > 0 {
		if res[x][y-1] == 0 {
			doUniquePaths(m, n, x, y-1, res)
		}
		r2 = res[x][y-1]
	}
	res[x][y] = r1 + r2
}
