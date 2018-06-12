package matrix

import (
	"fmt"
	"testing"
)

func TestCholesky(t *testing.T) {
	mat := LoadCSV("data/choltest.csv")
	_ = Cholesky(mat)
}

func TestIsSymmetric(t *testing.T) {
	q := Eye(4)
	//q.addScaledRow(1, 0, 1)
	fmt.Println(q.isSymmetric())
}
