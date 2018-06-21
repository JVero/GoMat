package matrix

import (
	"testing"
)

func TestCSVParse(t *testing.T) {
	_ = LoadCSV("data/sampledata.csv")
}

func TestCSVToStrings(t *testing.T) {
	a := LoadCSV("data/sampledata.csv")
	_ = dataToStrings(a)
}

func TestSaveMatrix(t *testing.T) {
	a := LoadCSV("data/sampledata.csv")
	a.ToCSV("data/outputfile.csv")
}

func TestCompression(t *testing.T) {
	a := LoadCSV("data/bigdata.csv")
	a.ToGZ("data/bigdata")
	a.ToCSV("data/bigdata.csv")
}

func TestGob(t *testing.T) {
	a := LoadCSV("data/sampledata.csv")
	a.ToGob("data/sampledata.gob")
}

func TestGobLoad(t *testing.T) {
	_ = LoadGob("data/sampledata.gob")
}


