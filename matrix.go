package matrix

import (
	"bytes"
	"strconv"
)

type matrix struct {
	width  int
	height int
	values [][]float64
}

// CreateMatrix creates an empty 2D matrix that is xDim, yDim
func CreateMatrix(rowHeight int, columnWidth int) matrix {
	matVals := make([][]float64, rowHeight)
	for row := range matVals {
		matVals[row] = make([]float64, columnWidth)
	}
	return matrix{rowHeight, columnWidth, matVals}
}

func (m *matrix) assignValue(row int, column int, val float64) {
	m.values[row][column] = val
}

// InitMatrix is the basic initializer that takes an arbitrary amount of int slices as rows
// and returns the matrix datatype
func InitMatrix(numRows int, numCols int, rows ...[]float64) matrix {
	if numRows != len(rows) {
		return matrix{}
	}
	for row_ind := range rows {
		if len(rows[row_ind]) != numCols {
			return matrix{}
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

func IdentityMatrix(size int) matrix {
	returnMatrix := CreateMatrix(size, size)
	for i := 0; i < size; i++ {
		returnMatrix.assignValue(i, i, 1)
	}
	return returnMatrix
}

func (m matrix) String() string {
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

func (m matrix) add(n matrix) matrix {
	if m.width != n.width || m.height != n.height {
		return matrix{}
	}
	retMat := CreateMatrix(m.width, m.height)
	for i := range m.values {
		for j := range m.values[i] {
			retMat.values[i][j] = m.values[i][j] + n.values[i][j]
		}
	}
	return retMat

}

func (m matrix) GetDims() (int, int) {
	return m.width, m.height
}
func (m matrix) sub(n matrix) matrix {
	if m.width != n.width || m.height != n.height {
		return matrix{}
	}
	retMat := CreateMatrix(m.width, m.height)
	for i := range m.values {
		for j := range m.values[i] {
			retMat.values[i][j] = m.values[i][j] - n.values[i][j]
		}
	}
	return retMat
}

func (m matrix) multiply(n matrix) matrix {
	var retMat matrix
	if len(m.values) == len(n.values[0]) { // C1 == R2 {
		retMat = CreateMatrix(len(m.values), len(n.values[0]))
	} else {
		return matrix{}
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
