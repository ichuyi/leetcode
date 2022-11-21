package main

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	x, y := 0, 0
	result := make([]int, 0)
	for x <= (m-1)/2 && y <= (n-1)/2 {
		result = append(result, circle(matrix, m, n, x, y)...)
		x++
		y++
	}
	return result
}

func circle(matrix [][]int, m, n, x, y int) []int {
	result := make([]int, 0)
	result = append(result, line(matrix, x, y, 0, 1, n-2*y)...)
	result = append(result, line(matrix, x+1, n-y-1, 1, 0, m-2*x-1)...)
	if m-2*x-1 > 0 {
		result = append(result, line(matrix, m-x-1, n-y-2, 0, -1, n-2*y-1)...)
	}
	if n-2*y-1 > 0 {
		result = append(result, line(matrix, m-x-2, y, -1, 0, m-x*2-2)...)
	}
	return result
}
func line(matrix [][]int, startX, startY, stepX, stepY, step int) []int {
	result := make([]int, 0)
	for step > 0 {
		result = append(result, matrix[startX][startY])
		startX += stepX
		startY += stepY
		step--
	}
	return result
}
