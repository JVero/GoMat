package matrix

// The purpose of this file is to facilitate io functions, such as loading or saving matrices
import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadMat(filename string) Matrix {
	iobuf, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer iobuf.Close()
	dat := csv.NewReader(iobuf)

	strings, err := dat.ReadAll()
	if err != nil {
		panic(err)
	}
	floats := csvToData(strings)
	numRows := len(floats)
	numCols := len(floats[0])
	return InitMatrix(numRows, numCols, floats...)
}

func csvToData(data [][]string) [][]float64 {
	returnValue := make([][]float64, len(data))

	for i := range data {
		returnValue[i] = make([]float64, len(data[i]))
		for j := range data[i] {
			val, err := strconv.ParseFloat(strings.Trim(data[i][j], " "), 64)

			if err != nil {
				fmt.Println(err.Error())
				returnValue[i][j] = 0
			} else {
				returnValue[i][j] = val
			}
		}
	}
	return returnValue
}


func dataToStrings(mat Matrix) [][]string{
    numRows, numCols := mat.GetDims()
    retVal := make([][]string, numRows)
    for row_ind, row := range(mat.values) {
        retVal[row_ind] = make([]string, numCols)
        for col_ind  := range(row) {
            retVal[row_ind][col_ind] = strconv.FormatFloat(mat.At(row_ind, col_ind), 'E', -1, 64)
        }
    }
    return retVal
}

func MatrixToCSV(mat Matrix, fileName string) error {
    strings := dataToStrings(mat)

    file, err := os.Create(fileName)
    if err != nil {
        return err
    }

    fileWriter := csv.NewWriter(file)
    err = fileWriter.WriteAll(strings)
    if err != nil {
        fmt.Println(err.Error())
        return err
    }

    return nil
}


