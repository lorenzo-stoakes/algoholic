package algoholic

func BinarySearchRecursive(ns []int, n int) int {
	var search func([]int, int, int) int
	search = func(ns []int, n, offset int) int {
		if len(ns) == 0 {
			return -1
		}

		half := len(ns) / 2
		val := ns[half]

		switch {
		case n == val:
			return half + offset
		case n < val:
			return search(ns[:half], n, offset)
		default:
			// n > val
			return search(ns[half+1:], n, offset+half+1)
		}
	}

	return search(ns, n, 0)
}

func BinarySearchIterative(ns []int, n int) int {
	if len(ns) == 0 {
		return -1
	}

	from, to := 0, len(ns)-1

	for {
		// Avoid overflow.
		mid := to - (to-from)/2
		val := ns[mid]

		switch {
		case n == val:
			return mid
		case n < val:
			to = mid - 1
		default:
			// n > val
			from = mid + 1
		}

		if from > to {
			return -1
		}
	}
}
