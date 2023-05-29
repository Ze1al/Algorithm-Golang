package offer

import (
	"fmt"
	"testing"
)

func Test_coinChange(t *testing.T) {
	coins := []int{1, 2, 5}

	res := coinChange(coins, 11)
	fmt.Println(res)
}
