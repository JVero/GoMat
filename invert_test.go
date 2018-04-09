package matrix

import (
	"fmt"
	"testing"
)

func TestScale(t *testing.T) {
	testMat := Eye(4)
	testMat.scaleRow(2.0, 1)
	fmt.Printf("%v", testMat)
}
func TestPivot(t *testing.T) {
	testMat := Eye(4)
	testMat.pivot(0, 1)
	fmt.Printf("%v", testMat)
}

func TestAddScaledRow(t *testing.T) {
	testMat := Eye(4)
	testMat.addScaledRow(0, 1, 20)
	fmt.Printf("%v", testMat)
}

func TestAugmented(t *testing.T) {
	testMat := Eye(4)
	aug := makeAugmentedMatrix(testMat)
	fmt.Printf("%v", aug)
}

func TestInvert(t *testing.T) {
	testMat := LoadCSV("inv.csv")
	fmt.Printf("%v", testMat)
	inv := testMat.invert()
	fmt.Println(testMat.multiply(inv))
}
