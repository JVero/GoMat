package matrix

import (
	"testing"
)

func TestScale(t *testing.T) {
	testMat := Eye(4)
	testMat.scaleRow(2.0, 1)
}
func TestPivot(t *testing.T) {
	testMat := Eye(4)
	testMat.pivot(0, 1)
}

func TestAddScaledRow(t *testing.T) {
	testMat := Eye(4)
	testMat.addScaledRow(0, 1, 20)
}

func TestAugmented(t *testing.T) {
	testMat := Eye(4)
	aug := makeAugmentedMatrix(testMat)
	ToCSV(aug, "augtest.csv")
}

func TestInvert(t *testing.T) {
	testMat := LoadCSV("bigdata.csv")
	inv := testMat.invert()
	ToCSV(inv, "biginverse.csv")
}
