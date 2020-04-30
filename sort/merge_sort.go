package sort

// return merge sort
func MergeSort(arr []int) []int {
	return mergeSort(0, len(arr)-1, arr)
}

func mergeSort(left, right int, arr []int) []int {
	//stop
	if right - left < 2 {
		return arr[left:right]
	}
	//process
	mid := left + (right-left)>>1
	var (
		l, r []int
	)
	// thrill down
	l = mergeSort(left, mid, arr)
	r = mergeSort(mid+1, right, arr)
	// clear states
	return merge(l, r)
}

func merge(l, r []int) []int {
	result := make([]int, 0, len(l)+len(r))
	for len(l) != 0 && len(r) != 0 {
		if l[0] <= r[0] {
			result = append(result, l[0])
			l = l[1:]
		} else {
			result = append(result, r[0])
			r = r[1:]
		}
	}
	if len(l) != 0 {
		result = append(result, l...)
	}
	if len(r) != 0 {
		result = append(result, r...)
	}
	return result
}
