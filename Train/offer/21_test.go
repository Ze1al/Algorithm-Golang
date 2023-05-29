package offer

import (
	"fmt"
	"testing"
)

func Test_removeElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3

	res := removeElement(nums, val)
	fmt.Println(res)
}
