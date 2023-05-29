package meeting

import (
	"fmt"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	n := 11
	res := hammingWeight(uint32(n))
	fmt.Println(res)
}
