package sort

import (
	"github.com/googege/gotools/public"
)
// todo for pro.
func QuickSort(arr []int) []int {
	quickSort(0, len(arr)-1, arr)
	return arr
}

func quickSort(left, right int, arr []int) {
	// stop
	if left > right {
		return
	}
	//process
	ge := findPriIndex(left, right, arr)
	// trill down
	quickSort(left, ge-1, arr)
	quickSort(ge+1, right, arr)
}

func findPriIndex(left, right int, arr []int) int {
	pri := left
	index := pri + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pri] {
			public.Swap(i, index, arr)
			index++
		}
	}
	public.Swap(pri, index-1, arr)
	return index - 1
}
