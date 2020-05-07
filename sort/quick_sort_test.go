package sort

import (
	"github.com/googege/gotools/test"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Log(QuickSort(test.TestDatasSliceInt(100)))
}

func BenchmarkQuickSort(b *testing.B) {
	data := test.TestDatasSliceInt(10000000)
	for i := 0; i < b.N; i++ {
		QuickSort(data)
	}
}
func BenchmarkSort(b *testing.B) {
	data := test.TestDatasSliceInt(10000000)
	for i := 0; i < b.N; i++ {
		sort.Ints(data)
	}
}
