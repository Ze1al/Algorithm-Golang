package meeting

import (
	"fmt"
	"testing"
)

func TestInitLRU(t *testing.T) {
	lru := InitLRU(2)

	val1 := lru.Get(1)
	fmt.Printf("val1=%d，应该返回-1\n", val1)

	lru.Put(1, 1)
	lru.Put(2, 2)

	val2 := lru.Get(1)
	fmt.Printf("val2=%d，应该返回1\n", val2)

	lru.Put(3, 3)
	val3 := lru.Get(2)
	fmt.Printf("val3=%d，应该返回-1\n", val3)

	val4 := lru.Get(3)
	fmt.Printf("val4=%d，应该返回3\n", val4)
}
