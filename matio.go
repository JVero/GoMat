package matrix

// The purpose of this file is to facilitate io functions, such as loading or saving matrices
import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// LoadCSV takes a file called filename and
// loads it as a matrix, returning that matrix
func LoadCSV(filename string) Matrix {
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
	return New(numRows, numCols, floats...)
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

func dataToStrings(mat Matrix) [][]string {
	numRows, numCols := mat.Dimensions()
	retVal := make([][]string, numRows)
	for rowInd, row := range mat.values {
		retVal[rowInd] = make([]string, numCols)
		for colInd := range row {
			retVal[rowInd][colInd] = strconv.FormatFloat(mat.At(rowInd, colInd), 'E', -1, 64)
		}
	}
	return retVal
}

// ToCSV saves the matrix as a .csv file
func ToCSV(mat Matrix, fileName string) error {
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
