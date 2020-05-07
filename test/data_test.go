package test

import (
	"testing"
)

func TestTestDatasSliceInt(t *testing.T) {
	t.Log(TestDatasSliceInt(10))
}
func BenchmarkTestDatasSliceInt(b *testing.B) {

	for i := 0; i < b.N; i++ {
		TestDatasSliceInt(20)
	}
}
