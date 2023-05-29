package channel

import "testing"

func TestRandomSelect(t *testing.T) {
	RandomSelect()
}

func TestDeadLock(t *testing.T) {
	DeadLock()
}

func TestSelectChan(t *testing.T) {
	SelectChan()
}
