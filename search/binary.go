package search

func BinarySearch(arr []int, num int) bool {
	lo := 0
	hi := len(arr)

	for lo < hi {
		m := lo + (hi-lo)/2
		v := arr[int(m)]

		if v == num {
			return true
		} else if v > num {
			hi = int(m)
		} else {
			lo = int(m) + 1
		}
	}

	return false
}
