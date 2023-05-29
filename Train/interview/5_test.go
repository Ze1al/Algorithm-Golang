package interview

import (
	"fmt"
	"testing"
)

func Test_removeKDigits(t *testing.T) {
	nums := 32406
	res := removeKDigits(nums, 1)
	fmt.Println(res)

	res = removeKDigits(nums, 2)
	fmt.Println(res)
}

func Test_removeKDigits2(t *testing.T) {
	nums := 98765
	res := removeKDigits(nums, 1)
	fmt.Println(res)

	res = removeKDigits(nums, 2)
	fmt.Println(res)
}
