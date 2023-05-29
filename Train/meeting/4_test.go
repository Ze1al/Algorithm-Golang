package meeting

import (
	"fmt"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	s := "25525511135"
	res := restoreIpAddresses(s)
	fmt.Println(res)
}
