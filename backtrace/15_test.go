package backTrace

import (
	"fmt"
	"testing"
)

func Test_b(t *testing.T) {
	fmt.Printf("return=%d", b())
}

func Test_c(t *testing.T) {
	fmt.Printf("return=%d", *c())
}
