package interview

import (
	"testing"
	"time"
)

// 按照顺序打印 dog cat fish

func TestDog(t *testing.T) {
	for i := 0; i < 10; i++ {
		go Dog()
		go Cat()
		go Fish()
	}

	fish <- struct{}{}

	time.Sleep(10 * time.Second)
}
