package offer

import (
	"fmt"
	"testing"
)

func Test_generateMatrix(t *testing.T) {
	n := 3

	res := generateMatrix(n)
	fmt.Println(res)
}
