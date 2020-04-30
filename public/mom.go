package public

import (
	"math/rand"
	"time"
)

var (
	Seed = time.Now().UnixNano()
)

// swap x and y  in arr.
func Swap(x, y int, arr []int) {
	arr[x], arr[y] = arr[y], arr[x]
}

// return a new rand.Rand
func NewRand() *rand.Rand {
	return rand.New(rand.NewSource(Seed))
}
