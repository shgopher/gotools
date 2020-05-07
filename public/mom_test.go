package public

import (
	"testing"
)

func TestSwap(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	Swap(1, 4, s)
	t.Log(s)
}
func TestNewRand(t *testing.T) {
	t.Log(NewRand().Intn(10))
}
