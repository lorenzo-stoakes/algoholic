package algoholic

func LinearSearch(ns []int, n int) int {
	for i, m := range ns {
		if m == n {
			return i
		}
	}

	return -1
}
