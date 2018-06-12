package matrix

// The purpose of this file is to facilitate io functions, such as loading or saving matrices
import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
	"io/ioutil"
	"compress/gzip"
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
	for rowInd := 0; rowInd < numRows; rowInd++ {
		retVal[rowInd] = make([]string, numCols)
		for colInd := 0; colInd < mat.numCols; colInd++ {
			val := mat.Get(rowInd, colInd)
			retVal[rowInd][colInd] = strconv.FormatFloat(val, 'E', -1, 64)
		}
	}
	return retVal
}

// ToCSV saves the matrix as a .csv file
func (mat Matrix) ToCSV(filename string) error {
	strings := dataToStrings(mat)

	file, err := os.Create(filename)
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

// ToGZ saves the matrix as a .csv.gz file
// first the function saves as a csv file, then
// compresses it with gzip
func (mat Matrix) ToGZ(filename string) error {
	_ = mat.ToCSV("tempfile.csv")
	f,_ := os.Open("tempfile.csv")

	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	name := filename + ".csv" + ".gz"
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	w := gzip.NewWriter(f)
	w.Write(content)
	w.Close()
	os.Remove("tempfile.csv")
	return nil
}

