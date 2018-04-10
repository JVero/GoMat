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
	testMat := LoadCSV("bigdata.csv")
	prod := testMat.Add(testMat)
	fmt.Printf("%v", prod.At(0, 0))
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
	testMat := LoadCSV("bigishdata.csv")
	prod := testMat.multiply(testMat)
	fmt.Printf("%v", prod.At(0, 0))
	ToCSV(prod, "bigMult.csv")
}

func TestMultDiffDims(t *testing.T) {
	a := New(9, 9, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]float64{7, 8, 9, 1, 2, 3, 4, 5, 6},
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]float64{7, 8, 9, 1, 2, 3, 4, 5, 6},
		[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{4, 5, 6, 7, 8, 9, 1, 2, 3},
		[]float64{7, 8, 9, 1, 2, 3, 4, 5, 6})
	b := New(9, 5, []float64{1, 2, 3, 4, 5},
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
	fmt.Printf("%v\n", a.At(1, 1))
}

func TestEye(t *testing.T) {
	a := Eye(5)
	fmt.Printf("%v", a)
}

func TestDet(t *testing.T) {
	a := Eye(5)
	a.scaleRow(2, 1)
	fmt.Printf("%v\n", a.Det())
}

func TestTranspose(t *testing.T) {
	mat := LoadCSV("sampledata.csv")
	fmt.Printf("%v", mat.T())
}

func BenchmarkBigMatMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testMat := LoadCSV("bigdata.csv")
		_ = testMat.multiply(testMat)
	}
}

func TestPad(t *testing.T) {
	testMat := LoadCSV("sampledata1.csv")
	fmt.Printf("%v", testMat)
	fmt.Printf("%v", Pad(testMat))
}

func TestPartition(t *testing.T) {
	testMat := LoadCSV("sampledata1.csv")
	fmt.Printf("%v\n", testMat)
	padded := Pad(testMat)
	a1, a2, a3, a4 := Partition(padded)
	fmt.Printf("%v\n%v\n%v\n%v\n", a1, a2, a3, a4)
}

func TestCompose(t *testing.T) {
	testMat := LoadCSV("sampledata1.csv")
	fmt.Printf("%v\n", testMat)
	padded := Pad(testMat)
	a1, a2, a3, a4 := Partition(padded)
	fmt.Printf("%v\n%v\n%v\n%v\n", a1, a2, a3, a4)
}

func TestStrassen(t *testing.T) {
	bigMat := LoadCSV("bigdata.csv")
	newFile := bigMat.Strassen(bigMat)
	ToCSV(newFile, "bigStrassen.csv")
}
