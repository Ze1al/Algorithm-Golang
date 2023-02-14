package backTrace

import (
	"fmt"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	res := solveNQueens(4)
	fmt.Printf("res=%v", res)
}
