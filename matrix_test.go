package matrix

import (
	"fmt"
	"testing"
)

var testValue interface{}

func TestInit(t *testing.T) {
	_ = Empty(3, 3)
}

func TestAssign(t *testing.T) {
	_ = New(3, 3, []float64{1, 2, 3},
		[]float64{4, 5, 6},
		[]float64{7, 8, 9})
}

func TestAdd(t *testing.T) {
	a := New(9, 9, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]float64{7, 8, 9, 1, 2, 3, 4, 5, 6},
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]float64{7, 8, 9, 1, 2, 3, 4, 5, 6},
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]float64{7, 8, 9, 1, 2, 3, 4, 5, 6})

	b := New(9, 9, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]float64{7, 8, 9, 1, 2, 3, 4, 5, 6},
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]float64{7, 8, 9, 1, 2, 3, 4, 5, 6},
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]float64{7, 8, 9, 1, 2, 3, 4, 5, 6})
	_ = a.Add(b)
}

func TestSub(t *testing.T) {
	a := New(3, 3, []float64{1, 2, 3},
		[]float64{4, 5, 6},
		[]float64{7, 8, 9})
	b := New(3, 3, []float64{1, 2, 3},
		[]float64{4, 5, 6},
		[]float64{7, 8, 9})
	_ = a.Sub(b)
}

func TestBigAdd(t *testing.T) {
	testMat := LoadCSV("data/bigdata.csv")
	println(testMat.numRows, testMat.numCols)
	_ = testMat.Add(testMat)
}

func TestMultsameDims(t *testing.T) {
	a := New(9, 9, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{4, 5, 6, 7, 8, 9, 1, 2, 3}, []float64{7, 8, 9, 1, 2, 3, 4, 5, 6}, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{4, 5, 6, 7, 8, 9, 1, 2, 3}, []float64{7, 8, 9, 1, 2, 3, 4, 5, 6}, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []float64{4, 5, 6, 7, 8, 9, 1, 2, 3}, []float64{7, 8, 9, 1, 2, 3, 4, 5, 6})

	b := New(9, 9,
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]float64{7, 8, 9, 1, 2, 3, 4, 5, 6},
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]float64{7, 8, 9, 1, 2, 3, 4, 5, 6},
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]float64{7, 8, 9, 1, 2, 3, 4, 5, 6})
	_ = a.multiply(b)
}

func TestBigMults(t *testing.T) {
	testMat := LoadCSV("data/bigishdata.csv")
	prod := testMat.multiply(testMat)
	prod.ToCSV("data/bigMult.csv")
}

func TestMultDiffDims(t *testing.T) {
	a := New(9, 9,
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]float64{7, 8, 9, 1, 2, 3, 4, 5, 6},
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]float64{7, 8, 9, 1, 2, 3, 4, 5, 6},
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]float64{7, 8, 9, 1, 2, 3, 4, 5, 6})
	b := New(9, 5,
		[]float64{1, 2, 3, 4, 5},
		[]float64{4, 5, 6, 7, 8},
		[]float64{7, 8, 9, 1, 2},
		[]float64{1, 2, 3, 4, 5},
		[]float64{4, 5, 6, 7, 8},
		[]float64{7, 8, 9, 1, 2},
		[]float64{1, 2, 3, 4, 5},
		[]float64{4, 5, 6, 7, 8},
		[]float64{7, 8, 9, 1, 2})
	a.multiply(b)
}

func TestGetValue(t *testing.T) {
	a := New(9, 9, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]float64{7, 8, 9, 1, 2, 3, 4, 5, 6},
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]float64{7, 8, 9, 1, 2, 3, 4, 5, 6},
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]float64{7, 8, 9, 1, 2, 3, 4, 5, 6})
	fmt.Printf("%v\n", a.Get(1, 1))
}

func TestDet(t *testing.T) {
	a := Eye(5)
	a.scaleRow(2, 1)
	fmt.Printf("Determinant test: the determinant is %v\n", a.Det())
}

func TestTranspose(t *testing.T) {
	mat := LoadCSV("data/sampledata.csv")
	_ = mat.T()
}

func BenchmarkBigMatMul(b *testing.B) {
	testMat := LoadCSV("data/bigdata.csv")
	for i := 0; i < b.N; i++ {
		_ = testMat.multiply(testMat)
	}
}

func BenchmarkStrassen(b *testing.B) {
	testMat := LoadCSV("data/bigdata.csv")
	for i := 0; i < b.N; i++ {
		_ = testMat.Strassen(testMat)
	}
}

func TestPad(t *testing.T) {
	testMat := LoadCSV("data/sampledata1.csv")
	_ = Pad(testMat)
}

func TestPartition(t *testing.T) {
	testMat := LoadCSV("data/sampledata1.csv")
	padded := Pad(testMat)
	_, _, _, _ = Partition(padded)
}

/*
func TestCompose(t *testing.T) {
	testMat := LoadCSV("data/sampledata1.csv")
	padded := Pad(testMat)
	_, _, _, _ = Partition(padded)
}
*
/*
func TestStrassen(t *testing.T) {
	bigMat := LoadCSV("data/bigdata.csv")
	newFile := bigMat.Strassen(bigMat)
	newFile.ToCSV("data/bigStrassen.csv")
}
*/
