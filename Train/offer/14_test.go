package offer

import (
	"fmt"
	"testing"
)

func Test_myPow(t *testing.T) {
	x := 2
	n := 10

	res := myPow(float64(x), n)
	fmt.Printf("res=%v", res)
}
