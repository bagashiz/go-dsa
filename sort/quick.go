package sort

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivotIdx := partition(arr, lo, hi)

	quickSort(arr, lo, pivotIdx-1)
	quickSort(arr, pivotIdx+1, hi)
}

func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if arr[i] < pivot {
			idx++
			arr[i], arr[idx] = arr[idx], arr[i]
		}
	}

	idx++
	arr[hi], arr[idx] = arr[idx], arr[hi]

	return idx
}
