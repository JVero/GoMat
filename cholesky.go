package matrix

import (
	"math"
	"reflect"
)

func (m Matrix) isSymmetric() bool {
	return reflect.DeepEqual(m, m.T())
}

func Cholesky(m Matrix) Matrix {
	if !m.isSymmetric() {
		panic("Matrix:  the matrix must be symmetric")
	}
	L := Empty(m.height, m.width)
	for i := 0; i < m.height; i++ {
		for j := 0; j <= i; j++ {
			if i == j {
				sum := 0.0
				for k := 0; k < j; k++ {
					sum += math.Pow(L.values[j][k], 2)
				}
				L.values[j][j] = math.Sqrt(m.At(j, j) - sum)
			} else {
				sum := 0.0
				for k := 0; k < j; k++ {
					sum += L.values[i][k] * L.values[j][k]
				}
				L.values[i][j] = (1 / L.values[j][j]) * (m.values[i][j] - sum)
			}
		}
	}
	return L
}
