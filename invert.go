package matrix

import "fmt"

func (m *Matrix) scaleRow(scale float64, row int) {
	if row >= m.height {
		panic("Matrix: Row out of range of the size of the matrix")
	}
	for i := range m.values[row] {
		m.values[row][i] *= scale
	}

}

func (m *Matrix) pivot(row1, row2 int) {
	for i := range m.values[row1] {
		m.values[row1][i], m.values[row2][i] = m.values[row2][i], m.values[row1][i]
	}
}

func (m *Matrix) addScaledRow(row1 int, row2 int, scale float64) {
	if row1 >= m.height || row2 >= m.height {
		panic("Matrix: Row outof range of the size of the matrix")
	}
	for i := range m.values[row1] {
		m.values[row1][i] += scale * m.values[row2][i]
	}
}

func makeAugmentedMatrix(m Matrix) Matrix {
	if m.width != m.height {
		panic("Matrix: matrix must be square to invert it")
	}
	augmentedMatrix := make([][]float64, m.width)
	for i := range augmentedMatrix {
		augmentedMatrix[i] = make([]float64, 2*m.width)
	}
	for i, row := range m.values {
		for j, val := range row {
			augmentedMatrix[i][j] = val
			if i == j {
				augmentedMatrix[i][j+m.width] = 1
			} else {
				augmentedMatrix[i][j+m.width] = 0
			}
		}
	}
	retVal := New(m.height, 2*m.width, augmentedMatrix...)
	return retVal
}

func (m Matrix) invert() Matrix {
	aug := makeAugmentedMatrix(m)
	for i := range aug.values {
		pivotRow := i
		for aug.At(i, i) == 0 && pivotRow < m.height {
			aug.pivot(i, pivotRow)
			pivotRow++
			fmt.Printf("%v", aug)
		}
		if aug.At(i, i) == 0 {
			panic("Matrix: Singular Matrix")
		}
		for j := range aug.values {
			if i == j {
				continue
			}
			scale := -1 * aug.At(j, i) / aug.At(i, i)
			fmt.Printf("Scale is %f\n for %f and %f\n", scale, aug.At(j, i), aug.At(i, i))
			aug.addScaledRow(j, i, scale)
			fmt.Printf("%v", aug)
		}
		for k := range aug.values {
			scale := aug.At(k, k)
			aug.scaleRow(1/scale, k)
			fmt.Printf("%v", aug)
		}

	}
	return extractFromAugment(aug)
}

func extractFromAugment(m Matrix) Matrix {
	retVal := make([][]float64, m.width)
	//numCols := m.width
	for i := range retVal {
		retVal[i] = m.values[i][m.height/2:]
		fmt.Println(retVal[i])
	}
	return New(m.height/2, m.width, retVal...)
}
