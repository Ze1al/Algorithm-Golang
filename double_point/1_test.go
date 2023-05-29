package double_point

import (
	"fmt"
	"testing"
)

func Test_removeElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	res := removeElement(nums, 3)
	fmt.Printf("res=%v", res)
}
