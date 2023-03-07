package monotonic_stack

import (
	"fmt"
	"testing"
)

func Test_trap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(height)
	fmt.Printf("res=%v", res)
}
