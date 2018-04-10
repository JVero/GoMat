package matrix

import (
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
	ToCSV(a, "outputfile.csv")
}
