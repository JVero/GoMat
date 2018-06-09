package matrix

import (
	"bytes"
	"strconv"
	"sync"
)

// Matrix is the default 2D data type that represents a matrix
type Matrix struct {
	numRows int
	numCols int
	values  []float64
}

// Empty creates an empty 2D Matrix that is rowHeight x columnWidth
func Empty(numRows int, numCols int) Matrix {
	matVals := make([]float64, numRows*numCols)
	return Matrix{numRows, numCols, matVals}
}

func (m *Matrix) Get(row int, column int) float64 {
	if row > m.numRows || column > m.numCols {
		panic("Out of range")
	}
	return m.values[m.numCols*row+column]
}

func (m *Matrix) assignValue(row int, column int, val float64) {
	if row > m.numRows || column > m.numCols {
		panic("Out of range")
	}
	m.values[m.numCols*row+column] = val
}

// New is the basic initializer that takes an arbitrary amount of int slices as rows
// and returns the Matrix datatype
func New(numRows int, numCols int, rows ...[]float64) Matrix {
	if numRows != len(rows) {
		panic("Matrix: The dimensions of the rows must match numRows")
	}
	for rowInd := range rows {
		if len(rows[rowInd]) != numCols {
			panic("Matrix: The length of all the columns must match numCols")
		}
	}
	retMat := Empty(numRows, numCols)
	for xval := 0; xval < retMat.numRows; xval++ {
		for yval := 0; yval < retMat.numCols; yval++ {
			retMat.assignValue(xval, yval, rows[xval][yval])
		}
	}
	return retMat
}

//Eye creates a size x size identity Matrix
func Eye(size int) Matrix {
	returnMatrix := Empty(size, size)
	for i := 0; i < size; i++ {
		returnMatrix.assignValue(i, i, 1)
	}
	return returnMatrix
}

// T is transpose
func (m Matrix) T() Matrix {
	numRows := m.numCols
	numCols := m.numRows
	retMatData := make([][]float64, numRows)
	for i := range retMatData {
		retMatData[i] = make([]float64, numCols)
		for j := range retMatData[i] {
			retMatData[i][j] = m.Get(j, i)
		}
	}
	return New(numRows, numCols, retMatData...)
}

func (m Matrix) String() string {
	var retString bytes.Buffer
	for i := 0; i < m.numRows; i++ {
		if i == 0 {
			retString.WriteString("[ ")
		} else {
			retString.WriteString("  ")
		}
		for j := 0; j < m.numCols; j++ {
			retString.WriteString(strconv.FormatFloat(m.Get(i, j), 'f', 2, 64))
			if j != m.numCols-1 {
				retString.WriteString(", ")
			}
		}
		if i == m.numRows-1 {
			retString.WriteString(" ]")
		}
		retString.WriteString("\n")
	}
	return retString.String()
}

// Add m and n
func (m Matrix) Add(n Matrix) Matrix {
	if m.numCols != n.numCols || m.numRows != n.numRows {
		panic("Matrix: Dimensions must match")
	}

	retMat := Empty(m.numCols, m.numRows)
	for i := 0; i < m.numRows; i++ {
		for j := 0; j < m.numCols; j++ {
			newVal := m.Get(i, j) + n.Get(i, j)
			retMat.assignValue(i, j, newVal)
		}
	}
	return retMat

}

// Dimensions is a trivial function that returns the numCols and the numRows of the matrix
func (m Matrix) Dimensions() (int, int) {
	return m.numRows, m.numCols
}

//Sub calculates m - n
func (m Matrix) Sub(n Matrix) Matrix {
	if m.numCols != n.numCols || m.numRows != n.numRows {
		panic("Matrix:  Dimensions must match")
	}
	retMat := Empty(m.numCols, m.numRows)
	for i := 0; i < m.numRows; i++ {
		for j := 0; j < m.numCols; j++ {
			val := m.Get(i, j) - n.Get(i, j)
			retMat.assignValue(i, j, val)
		}
	}
	return retMat
}

func (m Matrix) multiply(n Matrix) Matrix {
	var retMat Matrix
	var wg sync.WaitGroup
	if m.numCols == n.numRows { // C1 == R2 {
		retMat = Empty(m.numRows, n.numCols)
	} else {
		panic("Dimension mismatch")
	}
	for i := 0; i < m.numRows; i++ {
		wg.Add(1)
		go func(ig int) {
			defer wg.Done()
			for j := 0; j < n.numCols; j++ {
				for k := 0; k < m.numCols; k++ {
					if ig >= retMat.numRows || j >= retMat.numCols {
						panic("first")
					}
					if ig >= m.numRows || k >= m.numCols {
						panic("second")
					}
					if k >= n.numRows || j >= n.numCols {
						panic("third")
					}
					val1 := retMat.Get(ig, j)
					val2 := m.Get(ig, k)
					val3 := n.Get(k, j)
					val := val1 + val2*val3
					retMat.assignValue(ig, j, val)
				}
			}
		}(i)
	}
	wg.Wait()
	return retMat
}

// Det is the determinant
func (m Matrix) Det() (det float64) {
	var wg sync.WaitGroup
	det = 1
	for i := 0; i < m.numRows; i++ {
		pivotRow := i
		for m.Get(i, i) == 0 && pivotRow < m.numRows {
			m.pivot(i, pivotRow)
			pivotRow++
		}

		if m.Get(i, i) == 0 {
			return 0
		}
		for j := 0; j < m.numRows; j++ {
			if i == j {
				continue
			}
			wg.Add(1)
			go func(jg, ig int) {
				defer wg.Done()
				val1 := m.Get(jg, ig)
				val2 := m.Get(ig, ig)
				scale := -1 * val1 / val2
				m.addScaledRow(jg, ig, scale)
			}(j, i)
		}

	}
	for i := 0; i < m.numRows; i++ {
		det *= m.Get(i, i)
	}
	return det
}
