package public

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	Swap(1, 4, s)
	fmt.Println(s)
}
func TestNewRand(t *testing.T) {
	fmt.Println(NewRand().Intn(10))
}
