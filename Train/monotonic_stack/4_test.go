package monotonic_stack

import (
	"fmt"
	"testing"
)

func Test_trap(t *testing.T) {
	height := []int{4, 2, 0, 3, 2, 5}

	max := trap(height)
	fmt.Println(max)
}
