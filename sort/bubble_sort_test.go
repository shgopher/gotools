package sort

import (
	"github.com/googege/gotools/test"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	data := test.TestDatasSliceInt(10)
	t.Log(data)
	BubbleSort(data)
	t.Log(data)
}
func BenchmarkBubbleSort(b *testing.B) {
	data := test.TestDatasSliceInt(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(data)
	}
}
