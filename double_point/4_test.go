package double_point

import "testing"

func Test_backspaceCompare(t *testing.T) {
	s := "xywrrmp"
	s2 := "xywrrmu#p"

	backspaceCompare(s, s2)

}
