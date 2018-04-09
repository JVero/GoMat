package matrix

import (
	"fmt"
	"testing"
)

func TestCholesky(t *testing.T) {

}

func TestIsSymmetric(t *testing.T) {
	q := Eye(4)
	fmt.Println(q.isSymmetric())
}

func TestChol(t *testing.T) {
	mat := LoadCSV("choltest.csv")
	fmt.Printf("%v", mat)
	fmt.Printf("%v", cholesky(mat))
}
