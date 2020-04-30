package sort

import (
	"fmt"
	"github.com/googege/gotools/test"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	data := test.TestDatasSliceInt(100)
	fmt.Println(data)
	BubbleSort(data)
	fmt.Println(data)
}
func BenchmarkBubbleSort(b *testing.B) {
	data := test.TestDatasSliceInt(100)
	for i := 0; i < b.N; i++ {
		BubbleSort(data)
	}
}
