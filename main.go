package main

import "github.com/phanvudinh/practice-go/matrix"

func main() {
	m := matrix.InputMatrix()
	matrix.DisplayMatrix(m)

	// matrix.DisplayMatrix(matrix.Multi(m, m1))
	// matrix.DisplayMatrix(echelonMatrix)
	// combinedMatrix := matrix.GenerateOneMatrixByCombineUnitMatrix(m)
	// matrix.DisplayMatrix(combinedMatrix)

	echelonMatrix := matrix.EchelonMatrix(m)
	matrix.DisplayMatrix(echelonMatrix)

	// inversionMatrix := matrix.InversionMatrix(m)
	// matrix.DisplayMatrix(inversionMatrix)

	// matrix.DisplayMatrix(matrix.Multi(inversionMatrix, m))

}
