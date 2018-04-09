package matrix

import (
	"bytes"
	"strconv"
	"sync"
)

// Matrix is the default 2D data type that represents a matrix
type Matrix struct {
	width  int
	height int
	values [][]float64
}

// Empty creates an empty 2D Matrix that is rowHeight x columnWidth
func Empty(rowHeight int, columnWidth int) Matrix {
	matVals := make([][]float64, rowHeight)
	for row := range matVals {
		matVals[row] = make([]float64, columnWidth)
	}
	return Matrix{rowHeight, columnWidth, matVals}
}

func (m *Matrix) assignValue(row int, column int, val float64) {
	m.values[row][column] = val
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
	for xval, row := range rows {
		for yval, matval := range row {
			retMat.assignValue(xval, yval, matval)
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
	numRows := m.height
	numCols := m.width
	retMatData := make([][]float64, numRows)
	for i := range retMatData {
		retMatData[i] = make([]float64, numCols)
		for j := range retMatData[i] {
			println(i, j)
			_ = m.values[j][i]
			retMatData[i][j] = m.values[j][i]
		}
	}
	return New(numRows, numCols, retMatData...)
}

// At returns the value at row, col
func (m Matrix) At(row, col int) float64 {
	return (m.values[row][col])
}
func (m Matrix) String() string {
	var retString bytes.Buffer
	for i := range m.values {
		if i == 0 {
			retString.WriteString("[ ")
		} else {
			retString.WriteString("  ")
		}
		for j := range m.values[i] {
			retString.WriteString(strconv.FormatFloat(m.values[i][j], 'f', 2, 64))
			if j != len(m.values[i])-1 {
				retString.WriteString(", ")
			}
		}
		if i == len(m.values)-1 {
			retString.WriteString(" ]")
		}
		retString.WriteString("\n")
	}
	return retString.String()
}

// Add m and n
func (m Matrix) Add(n Matrix) Matrix {
	if m.width != n.width || m.height != n.height {
		panic("Matrix: Dimensions must match")
	}

	retMat := Empty(m.width, m.height)
	for i := range m.values {
		for j := range m.values[i] {
			retMat.values[i][j] = m.values[i][j] + n.values[i][j]
		}
	}
	return retMat

}

// Dimensions is a trivial function that returns the width and the height of the matrix
func (m Matrix) Dimensions() (int, int) {
	return m.width, m.height
}

//Sub calculates m - n
func (m Matrix) Sub(n Matrix) Matrix {
	if m.width != n.width || m.height != n.height {
		panic("Matrix:  Dimensions must match")
	}
	retMat := Empty(m.width, m.height)
	for i := range m.values {
		for j := range m.values[i] {
			retMat.values[i][j] = m.values[i][j] - n.values[i][j]
		}
	}
	return retMat
}

func (m Matrix) multiply(n Matrix) Matrix {
	var retMat Matrix
	var wg sync.WaitGroup
	if len(m.values) == len(n.values[0]) { // C1 == R2 {
		retMat = Empty(len(m.values), len(n.values[0]))
	} else {
		return Matrix{}
	}
	for i := range m.values {
		wg.Add(1)
		go func(ig int) {
			defer wg.Done()
			for j := range n.values[0] {
				for k := range m.values[0] {
					retMat.values[ig][j] += m.values[ig][k] * n.values[k][j]
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
	for i := range m.values {
		pivotRow := i
		for m.At(i, i) == 0 && pivotRow < m.height {
			m.pivot(i, pivotRow)
			pivotRow++
		}

		if m.At(i, i) == 0 {
			return 0
		}
		for j := range m.values {
			if i == j {
				continue
			}
			wg.Add(1)
			go func(jg, ig int) {
				defer wg.Done()
				scale := -1 * m.At(jg, ig) / m.At(ig, ig)
				m.addScaledRow(jg, ig, scale)
			}(j, i)
		}

	}
	for i := range m.values {
		det *= m.At(i, i)
	}
	return det
}
