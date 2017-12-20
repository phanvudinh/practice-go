package main

import "fmt"

func main() {
	matrix := inputMatrix()
	displayMatrix(matrix)
	detMatrix(matrix)
}

func inputMatrix() [][]int {
	var M, N int
	for M == 0 || N == 0 {
		fmt.Print("Row's length: ")
		fmt.Scanf("%d", &M)
		fmt.Printf("Column's length: ")
		fmt.Scanf("%d", &N)
	}

	var matrix = make([][]int, M, M)
	for i := 0; i < M; i++ {
		var row = make([]int, N, N)
		for j := 0; j < N; j++ {
			fmt.Printf("Enter matrix[%d][%d]: ", i, j)
			fmt.Scanf("%d", &row[j])
		}
		matrix[i] = row
	}

	return matrix
}

func displayMatrix(matrix [][]int) {
	fmt.Println("Matrix")
	M, N := directionMatrixSize(matrix)
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d ", matrix[i][j])
			if j == N-1 {
				fmt.Println("")
				break
			}
		}
	}
}

func detMatrix(matrix [][]int) {
	m, n := directionMatrixSize(matrix)
	if m != n {
		panic("m != n")
	}
	if m == 1 {
	} else {
		for i := 0; i < m; i++ {
			displayMatrix(subMatrix(matrix, 1, i))
		}
	}
}

func directionMatrixSize(matrix [][]int) (int, int) {
	return len(matrix), len(matrix[0])
}

func subMatrix(matrix [][]int, withoutRow int, withoutCol int) [][]int {
	m, _ := directionMatrixSize(matrix)
	var newMatrix = make([][]int, m-1, m-1)
	var k, l int
	for i := 0; i < m; i++ {
		if i == withoutRow-1 {
			continue
		}
		l = 0
		var rowNewMatrix = make([]int, m-1, m-1)
		for j := 0; j < m; j++ {
			if j == withoutCol-1 {
				continue
			}
			rowNewMatrix[l] = matrix[i][j]
			l++
		}
		k++
		newMatrix[k] = rowNewMatrix
	}

	return newMatrix
}
