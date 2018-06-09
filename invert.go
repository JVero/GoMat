package matrix

import "sync"

func (m *Matrix) scaleRow(scale float64, row int) {
	if row >= m.numRows {
		panic("Matrix: Row out of range of the size of the matrix")
	}
	for i := 0; i < m.numCols; i++ {
		m.assignValue(row, i, m.Get(row, i)*scale)
	}

}

func (m *Matrix) pivot(row1, row2 int) {
	for i := 0; i < m.numCols; i++ {
		tval1 := m.Get(row1, i)
		tval2 := m.Get(row2, i)
		m.assignValue(row1, i, tval2)
		m.assignValue(row2, i, tval1)
	}
}

func (m *Matrix) addScaledRow(row1 int, row2 int, scale float64) {
	if row1 >= m.numRows || row2 >= m.numRows {
		panic("Matrix: Row outof range of the size of the matrix")
	}
	for i := 0; i < m.numCols; i++ {
		val := m.Get(row1, i)
		scaledVal := scale * m.Get(row2, i)
		newVal := val + scaledVal
		m.assignValue(row1, i, newVal)
	}
}

func makeAugmentedMatrix(m Matrix) Matrix {
	if m.numCols != m.numRows {
		panic("Matrix: matrix must be square to invert it")
	}
	augmentedMatrix := make([][]float64, m.numCols)
	for i := range augmentedMatrix {
		augmentedMatrix[i] = make([]float64, 2*m.numCols)
	}

	for i := 0; i < m.numRows; i++ {
		for j := 0; j < m.numCols; j++ {
			val := m.Get(i, j)
			augmentedMatrix[i][j] = val
			if i == j {
				augmentedMatrix[i][j+m.numCols] = 1
			} else {
				augmentedMatrix[i][j+m.numCols] = 0
			}
		}
	}
	retVal := New(m.numRows, 2*m.numCols, augmentedMatrix...)
	return retVal
}

func (m Matrix) invert() Matrix {
	aug := makeAugmentedMatrix(m)
	var wg sync.WaitGroup
	for i := 0; i < aug.numCols/2; i++ {
		pivotRow := i
		for aug.Get(i, i) == 0 && pivotRow < m.numRows {
			aug.pivot(i, pivotRow)
			pivotRow++
		}
		if aug.Get(i, i) == 0 {
			panic("Matrix: Singular Matrix")
		}
		for j := 0; j < aug.numRows; j++ {
			if i == j {
				continue
			}
			wg.Add(1)
			go func(jg, ig int) {
				defer wg.Done()
				if jg >= 2000 {
					panic(jg)
				}
				scale := -1 * aug.Get(jg, ig)
				scale /= aug.Get(ig, ig)
				aug.addScaledRow(jg, ig, scale)
			}(j, i)
		}
		wg.Wait()
		for k := 0; k < aug.numRows; k++ {
			wg.Add(1)
			go func(kg int) {
				defer wg.Done()
				scale := aug.Get(kg, kg)
				aug.scaleRow(1/scale, kg)
			}(k)
		}
		wg.Wait()
	}
	return extractFromAugment(aug)
}

func extractFromAugment(m Matrix) Matrix {
	retVal := make([][]float64, m.numRows)
	row := make([]float64, m.numCols/2)
	//numCols := m.numCols
	for i := range retVal {
		for j := 0; j < m.numCols/2; j++ {
			row[j] = m.Get(i, m.numCols/2+j)
		}
		retVal[i] = row
	}
	return New(m.numRows, m.numCols/2, retVal...)
}
