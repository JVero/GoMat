package matrix

import (
	"bytes"
	"strconv"
)

type Matrix struct {
	width  int
	height int
	values [][]float64
}

// CreateMatrix creates an empty 2D Matrix that is xDim, yDim
func CreateMatrix(rowHeight int, columnWidth int) Matrix {
	matVals := make([][]float64, rowHeight)
	for row := range matVals {
		matVals[row] = make([]float64, columnWidth)
	}
	return Matrix{rowHeight, columnWidth, matVals}
}

func (m *Matrix) assignValue(row int, column int, val float64) {
	m.values[row][column] = val
}

// InitMatrix is the basic initializer that takes an arbitrary amount of int slices as rows
// and returns the Matrix datatype
func InitMatrix(numRows int, numCols int, rows ...[]float64) Matrix {
	if numRows != len(rows) {
		return Matrix{}
	}
	for rowInd := range rows {
		if len(rows[rowInd]) != numCols {
			return Matrix{}
		}
	}
	retMat := CreateMatrix(numRows, numCols)
	for xval, row := range rows {
		for yval, matval := range row {
			retMat.assignValue(xval, yval, matval)
		}
	}
	return retMat
}

//Eye creates a size x size identity Matrix
func Eye(size int) Matrix {
	returnMatrix := CreateMatrix(size, size)
	for i := 0; i < size; i++ {
		returnMatrix.assignValue(i, i, 1)
	}
	return returnMatrix
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
			retString.WriteString(strconv.FormatFloat(m.values[i][j], 'E', -1, 64))
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

func (m Matrix) add(n Matrix) Matrix {
	if m.width != n.width || m.height != n.height {
		return Matrix{}
	}
	retMat := CreateMatrix(m.width, m.height)
	for i := range m.values {
		for j := range m.values[i] {
			retMat.values[i][j] = m.values[i][j] + n.values[i][j]
		}
	}
	return retMat

}

func (m Matrix) GetDims() (int, int) {
	return m.width, m.height
}
func (m Matrix) sub(n Matrix) Matrix {
	if m.width != n.width || m.height != n.height {
		return Matrix{}
	}
	retMat := CreateMatrix(m.width, m.height)
	for i := range m.values {
		for j := range m.values[i] {
			retMat.values[i][j] = m.values[i][j] - n.values[i][j]
		}
	}
	return retMat
}

func (m Matrix) multiply(n Matrix) Matrix {
	var retMat Matrix
	if len(m.values) == len(n.values[0]) { // C1 == R2 {
		retMat = CreateMatrix(len(m.values), len(n.values[0]))
	} else {
		return Matrix{}
	}
	for i := range m.values {
		for j := range n.values[0] {
			for k := range m.values[0] {
				retMat.values[i][j] += m.values[i][k] * n.values[k][j]
			}
		}
	}
	return retMat
}
