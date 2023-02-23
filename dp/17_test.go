package dp

import (
	"fmt"
	"testing"
)

func Test_coinChange(t *testing.T) {
	coin := []int{1, 2, 5}
	amount := 11

	res := coinChange(coin, amount)
	fmt.Printf("res=%d", res)
}
