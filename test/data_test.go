package test

import (
	"fmt"
	"testing"
)

func TestTestDatasSliceInt(t *testing.T) {
	fmt.Println(TestDatasSliceInt(10))
}
func BenchmarkTestDatasSliceInt(b *testing.B) {

	for i := 0; i < b.N; i++ {
		TestDatasSliceInt(20)
	}
}
