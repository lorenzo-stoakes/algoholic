package algoholic

func generateRange(from, to int) []int {
	length := to - from + 1

	if length < 0 {
		panic("Invalid from, to.")
	}

	ret := make([]int, length)

	for i := 0; i < length; i++ {
		ret[i] = from + i
	}

	return ret
}
