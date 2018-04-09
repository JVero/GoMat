package matrix

import (
	"math"
	"reflect"
)

func (m Matrix) isSymmetric() bool {
	return reflect.DeepEqual(m, m.T())
}

func cholesky(m Matrix) Matrix {
	if !m.isSymmetric() {
		panic("Matrix:  the matrix must be symmetric")
	}
	L := Empty(m.height, m.width)
	println(m.height, m.width, "HW")
	for i := 0; i < m.height; i++ {
		for j := 0; j <= i; j++ {
			if i == j {
				sum := 0.0
				for k := 0; k < j; k++ {
					println(i, j, k)
					sum += math.Pow(L.values[j][k], 2)
				}
				println(m.At(j, j))
				L.values[j][j] = math.Sqrt(m.At(j, j) - sum)
			} else {
				println("else")
				sum := 0.0
				for k := 0; k < j; k++ {
					sum += L.values[i][k] * L.values[j][k]
				}
				println(i, j)
				L.values[i][j] = (1 / L.values[j][j]) * (m.values[i][j] - sum)
			}
		}
	}
	return L
}
