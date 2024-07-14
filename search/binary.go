package search

func BinarySearch(arr []int, num int) bool {
	lo := 0
	hi := len(arr)

	for lo < hi {
		m := lo + (hi-lo)/2
		v := arr[m]

		if v == num {
			return true
		} else if v > num {
			hi = m
		} else {
			lo = m + 1
		}
	}

	return false
}
