package listnode

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	LRU := Constructor(2)

	LRU.Put(1, 1)
	LRU.Put(2, 2)
	LRU.Get(1)
	LRU.Put(3, 3)
	LRU.Get(2)
	LRU.Put(4, 4)
	LRU.Get(1)
	LRU.Get(3)
	LRU.Get(4)
}
