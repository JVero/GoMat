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
    b := dataToStrings(a)
    fmt.Printf("%v", b)
}
