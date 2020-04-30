package sort

import "github.com/googege/gotools/public"
// when n < 1000, this is better.
func BubbleSort(arr []int) {

	for i := 0; i < len(arr); i++ {
		count := 0
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j+1] < arr[j] {
				count++
				public.Swap(j, j+1, arr)
			}
		}
		if count == 0 {
			break
		}
	}
}
