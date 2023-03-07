package greed

import (
	"fmt"
	"testing"
)

func Test_lemonadeChange(t *testing.T) {
	bill := []int{5, 5, 5, 10, 20}
	res := lemonadeChange(bill)
	fmt.Printf("res=%v", res)
}
