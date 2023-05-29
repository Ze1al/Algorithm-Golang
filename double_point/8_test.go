package double_point

import (
	"fmt"
	"testing"
)

func Test_reverseStr(t *testing.T) {
	s := "abcdefg"

	res := reverseStr(s, 2)
	fmt.Printf("res=%v", res)
}
