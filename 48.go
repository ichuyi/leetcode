package main

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		start, end := i, n-1-i
		for j := 1; j < end-start; j++ {
			t := matrix[start][j+start]
			matrix[start][j+start] = matrix[end-j][start]
			matrix[end-j][start] = matrix[end][end-j]
			matrix[end][end-j] = matrix[start+j][end]
			matrix[j+start][end] = t
		}
		matrix[start][start], matrix[start][end], matrix[end][end], matrix[end][start] = matrix[end][start], matrix[start][start], matrix[start][end], matrix[end][end]
	}
}
