package matrix

import (
	"fmt"
	"testing"
)

func TestCSVParse(t *testing.T) {
	a := LoadMat("sampledata.csv")
	fmt.Println(a)
}
