package interview

import (
	"testing"
	"time"
)

func TestPrintNum(t *testing.T) {
	go PrintNum()
	go PrintLetter()

	letterChan <- struct{}{}

	time.Sleep(10 * time.Second)
}
