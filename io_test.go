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
	a.ToCSV("outputfile.csv")
}

func TestCompression(t *testing.T) {
	a := LoadCSV("bigdata.csv")
	a.ToGZ("bigdata")
	a.ToCSV("bigdata.csv")
}
