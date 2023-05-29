package offer

import (
	"fmt"
	"testing"
)

func Test_addBinary(t *testing.T) {
	a := "1010"
	b := "1011"

	res := addBinary(a, b)
	fmt.Println(res)
}
