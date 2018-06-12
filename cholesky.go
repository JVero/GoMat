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
	L := Empty(m.numRows, m.numCols)
	for i := 0; i < m.numRows; i++ {
		for j := 0; j <= i; j++ {
			if i == j {
				sum := 0.0
				for k := 0; k < j; k++ {
					sum += math.Pow(L.Get(j, k), 2)
				}
				L.assignValue(i, j, math.Sqrt(m.Get(j, j)-sum))
			} else {
				sum := 0.0
				for k := 0; k < j; k++ {
					sum += L.Get(i, j) * L.Get(j, k)
				}
				L.assignValue(i, j, (1 / L.Get(j, j) * (m.Get(i, j) - sum)))
			}
		}
	}
	return L
}
