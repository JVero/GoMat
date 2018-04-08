package matrix

import (
	"fmt"
	"testing"
)

func TestCSVParse(t *testing.T) {
	_ = LoadCSV("sampledata.csv")
}

func TestCSVToStrings(t *testing.T) {
	a := LoadCSV("sampledata.csv")
	_ = dataToStrings(a)
}

func TestSaveMatrix(t *testing.T) {
	a := LoadCSV("sampledata.csv")
	b := ToCSV(a, "outputfile.csv")
	fmt.Printf("%v", b)
}
