package matrix

import "math"

func Pad(m Matrix) Matrix {
	newDim := int(math.Pow(2, float64(int(math.Log2(float64(m.width-1))+1))))
	retMat := Empty(newDim, newDim)
	for i := range retMat.values {
		for j := range retMat.values {
			if i < m.width && j < m.width {
				retMat.values[i][j] = m.values[i][j]
			} else {
				retMat.values[i][j] = 0
			}
		}
	}
	return retMat

}

// Partition breaks up a padded matrix into four
func Partition(m Matrix) (a1, a2, a3, a4 Matrix) {
	if m.width > 1030 {
		panic("too big")
	}
	a1 = Empty(m.width/2, m.width/2)
	a2 = Empty(m.width/2, m.width/2)
	a3 = Empty(m.width/2, m.width/2)
	a4 = Empty(m.width/2, m.width/2)
	for i := range m.values {
		for j := range m.values[i] {
			if i < m.height/2 && j < m.width/2 {
				a1.values[i][j] = m.At(i, j)
			} else if i < m.height/2 && j >= m.width/2 {
				a2.values[i][j-m.width/2] = m.At(i, j)
			} else if i >= m.height/2 && j < m.width/2 {
				a3.values[i-m.height/2][j] = m.At(i, j)
			} else if i >= m.height/2 && j >= m.width/2 {
				a4.values[i-m.height/2][j-m.width/2] = m.At(i, j)
			}
		}
	}
	return a1, a2, a3, a4
}

func (m1 Matrix) Strassen(m2 Matrix) Matrix {
	if m1.height <= 256 {
		return m1.multiply(m2)
	}
	m1Padded := Pad(m1)
	m2Padded := Pad(m2)
	a1, a2, a3, a4 := Partition(m1Padded)
	b1, b2, b3, b4 := Partition(m2Padded)
	M1 := a1.Add(a4).Strassen(b1.Add(b4))
	M2 := a3.Add(a4).Strassen(b1)
	M3 := a1.Strassen(b2.Sub(b4))
	M4 := a4.Strassen(b3.Sub(b1))
	M5 := a1.Add(a2).Strassen(b4)
	M6 := a3.Sub(a1).Strassen(b1.Add(b2))
	M7 := a2.Sub(a4).Strassen(b3.Add(b4))
	c11 := M1.Add(M4).Sub(M5).Add(M7)
	c12 := M3.Add(M5)
	c21 := M2.Add(M4)
	c22 := M1.Sub(M2).Add(M3).Add(M6)
	return Compose(c11, c12, c21, c22, m1.height, m2.width)
}

func Compose(c1, c2, c3, c4 Matrix, originalHeight, originalWidth int) Matrix {
	m := Empty(originalHeight, originalHeight)
	for i := range m.values {
		for j := range m.values[i] {
			if i < c1.height && j < c1.width {
				if i >= len(c1.values) || j >= len(c1.values[0]) {
					panic("ahhhh")
				}
				m.values[i][j] = c1.values[i][j]
			} else if i < c1.height && j >= c1.width {
				m.values[i][j] = c2.values[i][j-c1.width]
			} else if i >= c1.height && j < c1.width {
				m.values[i][j] = c3.values[i-c1.height][j]
			} else if i >= c1.height && j >= c1.width {
				m.values[i][j] = c4.values[i-c1.height][j-c1.width]
			}
		}
	}

	return m
}
