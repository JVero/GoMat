package matrix

import (
	"math"
	"sync"
)

func Pad(m Matrix) Matrix {
	newDim := int(math.Pow(2, float64(int(math.Log2(float64(m.numCols-1))+1))))
	retMat := Empty(newDim, newDim)
	for i := 0; i < retMat.numRows; i++{
		for j := 0; j <  retMat.numCols; j++{
			if i < m.numCols && j < m.numCols {
				retMat.assignValue(i, j, m.Get(i, j))
			} else {
				retMat.assignValue(i, j, 0)
			}
		}
	}
	return retMat

}

// Partition breaks up a padded matrix into four
func Partition(m Matrix) (a1, a2, a3, a4 Matrix) {
	if m.numCols > 1030 {
		panic("too big")
	}
	a1 = Empty(m.numCols/2, m.numCols/2)
	a2 = Empty(m.numCols/2, m.numCols/2)
	a3 = Empty(m.numCols/2, m.numCols/2)
	a4 = Empty(m.numCols/2, m.numCols/2)
	for ig := 0; ig < m.numRows; ig++{
		go func(i int) {
			for j := 0; j < m.numCols; j++ {
				if i < m.numRows/2 && j < m.numCols/2 {
					a1.assignValue(i, j, m.Get(i, j))
				} else if i < m.numRows/2 && j >= m.numCols/2 {
					a2.assignValue(i, j-m.numCols/2, m.Get(i, j))
				} else if i >= m.numRows/2 && j < m.numCols/2 {
					a3.assignValue(i-m.numRows/2, j, m.Get(i, j))
				} else if i >= m.numRows/2 && j >= m.numCols/2 {
					a4.assignValue(i-m.numRows/2, j-m.numCols/2, m.Get(i, j))
				}
			}
		}(ig)
	}
	return a1, a2, a3, a4
}

func (m1 Matrix) Strassen(m2 Matrix) Matrix {
	if m1.numRows <= 256 {
		return m1.multiply(m2)
	}
	var wg sync.WaitGroup
	m1Padded := Pad(m1)
	m2Padded := Pad(m2)
	a1, a2, a3, a4 := Partition(m1Padded)
	b1, b2, b3, b4 := Partition(m2Padded)
	var M1, M2, M3, M4, M5, M6, M7 Matrix
	wg.Add(7)
	go func() {
		defer wg.Done()
		M1 = a1.Add(a4).Strassen(b1.Add(b4))
	}()
	go func() {
		defer wg.Done()
		M2 = a3.Add(a4).Strassen(b1)
	}()
	go func() {
		defer wg.Done()
		M3 = a1.Strassen(b2.Sub(b4))
	}()
	go func() {
		defer wg.Done()
		M4 = a4.Strassen(b3.Sub(b1))
	}()
	go func() {
		defer wg.Done()
		M5 = a1.Add(a2).Strassen(b4)
	}()
	go func() {
		defer wg.Done()
		M6 = a3.Sub(a1).Strassen(b1.Add(b2))
	}()
	go func() {
		defer wg.Done()
		M7 = a2.Sub(a4).Strassen(b3.Add(b4))
	}()
	wg.Wait()
	c11 := M1.Add(M4).Sub(M5).Add(M7)
	c12 := M3.Add(M5)
	c21 := M2.Add(M4)
	c22 := M1.Sub(M2).Add(M3).Add(M6)
	return Compose(c11, c12, c21, c22, m1.numRows, m2.numCols)
}

func Compose(c1, c2, c3, c4 Matrix, originalHeight, originalWidth int) Matrix {
	m := Empty(originalHeight, originalHeight)
	var wg sync.WaitGroup
	for i := 0; i < m.numRows; i++ {
		wg.Add(1)
		go func(ig int) {
			defer wg.Done()
			for j := 0; j < m.numCols; j++ {
				if ig < c1.numRows && j < c1.numCols {
					if ig >= c1.numRows || j >= c1.numCols {
						panic("ahhhh")
					}
					m.assignValue(ig ,j, c1.Get(ig, j))
				} else if ig < c1.numRows && j >= c1.numCols {
					m.assignValue(ig, j, c2.Get(ig, j-c1.numCols))
				} else if ig >= c1.numRows && j < c1. numCols {
					m.assignValue(ig, j, c3.Get(ig-c1.numRows, j))
				} else if ig >= c1.numRows && j >= c1.numCols {
					m.assignValue(ig, j, c4.Get(ig-c1.numRows, j-c1.numCols))
				}
			}
		}(i)
	}
	wg.Wait()
	return m
}
