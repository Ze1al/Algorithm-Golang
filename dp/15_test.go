package dp

import (
	"fmt"
	"testing"
)

func Test_change(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 5

	res := change(amount, coins)
	fmt.Printf("res=%d", res)
}
