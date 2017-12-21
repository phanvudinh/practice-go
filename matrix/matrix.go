package matrix

import "fmt"
import "math"
import "os"

// InputMatrix allow  enter matrix's value
func InputMatrix() [][]float64 {
	var M, N int
	for M == 0 || N == 0 {
		fmt.Print("Row's length: ")
		fmt.Scanf("%d", &M)
		fmt.Printf("Column's length: ")
		fmt.Scanf("%d", &N)
	}

	var matrix = make([][]float64, M, M)
	for i := 0; i < M; i++ {
		var row = make([]float64, N, N)
		for j := 0; j < N; j++ {
			fmt.Printf("Enter matrix[%d][%d]: ", i, j)
			fmt.Scanf("%f", &row[j])
		}
		matrix[i] = row
	}

	return matrix
}

// DisplayMatrix used to display matrix
func DisplayMatrix(matrix [][]float64) {
	fmt.Println("Matrix")
	M, N := DirectionMatrixSize(matrix)
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%f ", matrix[i][j])
			if j == N-1 {
				fmt.Println("")
				break
			}
		}
	}
}

// DetMatrix used to calculate DET MATRIX
func DetMatrix(matrix [][]float64) float64 {
	m, _ := DirectionMatrixSize(matrix)
	var det float64
	if m == 1 {
		det = matrix[0][0]
	} else {
		for i := 0; i < m; i++ {
			det += math.Pow(float64(-1), float64(1+i+1)) * matrix[0][i] * DetMatrix(SubMatrix(matrix, 0, i))
		}
	}
	return det
}

//DirectionMatrixSize present for direction matrix size
func DirectionMatrixSize(matrix [][]float64) (int, int) {
	return len(matrix), len(matrix[0])
}

// SubMatrix present for Sub-matrix
func SubMatrix(matrix [][]float64, withoutRow int, withoutCol int) [][]float64 {
	m, _ := DirectionMatrixSize(matrix)
	var newMatrix = make([][]float64, m-1, m-1)
	var k, l int
	for i := 0; i < m; i++ {
		if i == withoutRow {
			continue
		}
		l = 0
		var rowNewMatrix = make([]float64, m-1, m-1)
		for j := 0; j < m; j++ {
			if j == withoutCol {
				continue
			}
			rowNewMatrix[l] = matrix[i][j]
			l++
		}
		newMatrix[k] = rowNewMatrix
		k++
	}

	return newMatrix
}

//EchelonMatrix generate matrix become echelon matrix
func EchelonMatrix(matrix [][]float64) [][]float64 {
	m, n := DirectionMatrixSize(matrix)
	echelonMatrix := CopyMatrix(matrix)
	for i := 0; i < m-1; i++ {
		if echelonMatrix[i][i] == 0 {
			ExchangeRows(echelonMatrix, i)
		}
		for r := i + 1; r < m; r++ {
			delta := echelonMatrix[r][i] / echelonMatrix[i][i]
			for c := i; c < n; c++ {
				echelonMatrix[r][c] = echelonMatrix[r][c] - delta*echelonMatrix[i][c]
			}
		}
	}

	return echelonMatrix
}

//ExchangeRows exchange rows
func ExchangeRows(matrix [][]float64, rowIndex int) {
	m, n := DirectionMatrixSize(matrix)
	for i := rowIndex + 1; i < m; i++ {
		if matrix[i][rowIndex] == 1 {
			for j := 0; j < n; j++ {
				matrix[rowIndex][j], matrix[i][j] = matrix[i][j], matrix[rowIndex][j]
			}
			break
		}

		if matrix[i][rowIndex] != 0 {
			for j := 0; j < n; j++ {
				matrix[rowIndex][j], matrix[i][j] = matrix[i][j], matrix[rowIndex][j]
			}
		}
	}
}

//CopyMatrix function used to copy one matrix's value into other
func CopyMatrix(src [][]float64) [][]float64 {
	m, n := DirectionMatrixSize(src)
	tempMatrix := make([][]float64, m, m)
	for i := 0; i < m; i++ {
		row := make([]float64, n, n)
		copy(row, src[i])
		tempMatrix[i] = row
	}

	return tempMatrix
}

//IsUnitMatrix function is checked to consider it's unit matrix or not
func IsUnitMatrix(matrix [][]float64) bool {
	m, _ := DirectionMatrixSize(matrix)
	for i := 0; i < m; i++ {
		if matrix[i][i] != 1 {
			return false
		}
	}

	return true
}

//GererateUnitMatrix function used to generate new unit matrix
func GererateUnitMatrix(m int) [][]float64 {
	unitMatrix := make([][]float64, m, m)
	for i := 0; i < m; i++ {
		row := make([]float64, m, m)
		row[i] = 1
		unitMatrix[i] = row
	}

	return unitMatrix
}

//GenerateOneMatrixByCombineUnitMatrix function to
//generate another matrix by combine source matrix with unit matrix
func GenerateOneMatrixByCombineUnitMatrix(sourceMatrix [][]float64) [][]float64 {
	m, _ := DirectionMatrixSize(sourceMatrix)
	unitMatrix := GererateUnitMatrix(m)
	tempMatrix := make([][]float64, m, m)
	for i := 0; i < m; i++ {
		row := make([]float64, 2*m, 2*m)
		copy(row, sourceMatrix[i])
		copy(row[m:], unitMatrix[i])
		tempMatrix[i] = row
	}

	return tempMatrix
}

//InversionMatrix function to get inversion matrix from (A|I)
func InversionMatrix(sourceMatrix [][]float64) [][]float64 {
	m, _ := DirectionMatrixSize(sourceMatrix)
	det := DetMatrix(sourceMatrix)
	transposaedMatrix := TransposaedMatrix(sourceMatrix)
	tempMatrix := make([][]float64, m, m)
	for i := 0; i < m; i++ {
		row := make([]float64, m, m)
		for j := 0; j < m; j++ {
			row[j] = math.Pow(float64(-1), float64(j+i+2)) * 1 / det * DetMatrix(SubMatrix(transposaedMatrix, i, j))
		}
		tempMatrix[i] = row
	}

	return tempMatrix
}

//Multi function used to calculate operand multi two matrix
func Multi(matrixA, matrixB [][]float64) [][]float64 {
	m1, n1 := DirectionMatrixSize(matrixA)
	m2, n2 := DirectionMatrixSize(matrixB)
	if m2 != n1 {
		os.Exit(1)
	}
	tempMatrix := make([][]float64, m1, m1)
	for i := 0; i < m1; i++ {
		row := make([]float64, n2, n2)
		for j := 0; j < n2; j++ {
			var total float64
			for k := 0; k < n1; k++ {
				total += matrixA[i][k] * matrixB[k][j]
			}
			row[j] = total
		}
		tempMatrix[i] = row
	}

	return tempMatrix
}

// TransposaedMatrix function migrate source matrix become to transposaed matrix
func TransposaedMatrix(matrix [][]float64) [][]float64 {
	m, n := DirectionMatrixSize(matrix)
	temp := make([][]float64, n, n)
	for i := 0; i < m; i++ {
		row := make([]float64, m, m)
		for j := 0; j < n; j++ {
			row[j] = matrix[j][i]
		}
		temp[i] = row
	}
	return temp
}
