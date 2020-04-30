package sort

import (
	"fmt"
	"github.com/googege/gotools/test"
	"testing"
)

func TestMergeSort(t *testing.T) {
	data := test.TestDatasSliceInt(10)
	fmt.Println(MergeSort(data))
}
