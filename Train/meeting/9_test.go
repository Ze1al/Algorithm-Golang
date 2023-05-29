package meeting

import (
	"fmt"
	"testing"
)

func Test_convColumnNum(t *testing.T) {
	num := 701

	res := convColumnNum(num)
	fmt.Println(res)
}
