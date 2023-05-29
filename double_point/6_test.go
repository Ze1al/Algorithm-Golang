package double_point

import (
	"fmt"
	"testing"
)

func Test_replaceSpace(t *testing.T) {
	s := "We are happy."
	res := replaceSpace(s)
	fmt.Printf("res=%v", res)
}
