package matrix

import (
	"testing"
    "fmt"
)

func TestCSVParse(t *testing.T) {
	_ = LoadMat("sampledata.csv")
}

func TestCSVToStrings(t *testing.T) {
    a := LoadMat("sampledata.csv")
    _ = dataToStrings(a)
}

func TestSaveMatrix(t *testing.T) {
    a := LoadMat("sampledata.csv")
    b := MatrixToCSV(a, "outputfile.csv")
    fmt.Printf("%v", b)
}
