package main

func numIslands(grid [][]byte) int {
	l := len(grid)
	h := len(grid[0])
	f, i, j := getHead(grid, l, h)
	result := 0
	for f {
		dfsIsland(grid, l, h, i, j)
		result++
		f, i, j = getHead(grid, l, h)
	}
	return result
}
func dfsIsland(input [][]byte, l, h, i, j int) {
	if i >= l || j >= h || i < 0 || j < 0 {
		return
	}
	if input[i][j] == '0' {
		return
	}
	input[i][j] = '0'
	dfsIsland(input, l, h, i+1, j)
	dfsIsland(input, l, h, i, j+1)
	dfsIsland(input, l, h, i-1, j)
	dfsIsland(input, l, h, i, j-1)
}
func getHead(input [][]byte, l, h int) (bool, int, int) {
	for i := 0; i < l; i++ {
		for j := 0; j < h; j++ {
			if input[i][j] == '1' {
				return true, i, j
			}
		}
	}
	return false, 0, 0
}
