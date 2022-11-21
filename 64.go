package main

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	result := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		tmp := make([]int, 0, n)
		for j := 0; j < n; j++ {
			tmp = append(tmp, -1)
		}
		result = append(result, tmp)
	}
	return minPathSumDo(grid, m-1, n-1, result)
}
func minPathSumDo(grid [][]int, x, y int, result [][]int) int {
	if x+y == 0 {
		return grid[0][0]
	}
	x1, y1 := x, y-1
	x2, y2 := x-1, y
	if y1 >= 0 && result[x1][y1] == -1 {
		result[x1][y1] = minPathSumDo(grid, x1, y1, result)
	}
	if x2 >= 0 && result[x2][y2] == -1 {
		result[x2][y2] = minPathSumDo(grid, x2, y2, result)
	}
	if y1 < 0 {
		return result[x2][y2] + grid[x][y]
	}
	if x2 < 0 {
		return result[x1][y1] + grid[x][y]
	}
	r1 := min(result[x1][y1], result[x2][y2])
	return r1 + grid[x][y]
}
