package sort

import (
	"github.com/googege/gotools/test"
	"testing"
)

func TestMergeSort(t *testing.T) {
	data := test.TestDatasSliceInt(10)
	t.Log(MergeSort(data))
}
func BenchmarkMergeSort(b *testing.B) {
	data := test.TestDatasSliceInt(10000000)
	for i := 0; i < b.N; i++ {
		MergeSort(data)
	}
}
