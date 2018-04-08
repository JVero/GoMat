package matrix

import (
	"bytes"
	"strconv"
)

type matrix struct {
	width  int
	height int
	values [][]int
}

// CreateMatrix creates an empty 2D matrix that is xDim, yDim
func CreateMatrix(rowHeight int, columnWidth int) matrix {
	matVals := make([][]int, rowHeight)
	for row := range matVals {
		matVals[row] = make([]int, columnWidth)
	}
	return matrix{rowHeight, columnWidth, matVals}
}

func (m *matrix) assignValue(row int, column int, val int) {
	m.values[row][column] = val
}

// InitMatrix is the basic initializer that takes an arbitrary amount of int slices as rows
// and returns the matrix datatype
func InitMatrix(numRows int, numCols int, rows ...[]int) matrix {
	if numRows != len(rows) {
		return matrix{}
	}
	for row_ind := range rows {
		if len(rows[row_ind]) != numCols {
			return matrix{}
		}
	}
	retMat := CreateMatrix(numRows, numCols)
	a, b := retMat.GetDims()
	println(a)
	println(b)
	for xval, row := range rows {
		for yval, matval := range row {
			println(xval, yval)
			retMat.assignValue(xval, yval, matval)
		}
	}
	return retMat
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
			retString.WriteString(strconv.Itoa(m.values[i][j]))
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
