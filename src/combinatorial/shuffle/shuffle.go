package algoholic

import (
	"math/rand"
	"sort"
	"time"
)

type RandomOrderSlice []int

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (slice RandomOrderSlice) Len() int {
	return len(slice)
}

func (slice RandomOrderSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice RandomOrderSlice) Less(i, j int) bool {
	return rand.Float64() < 0.5
}

// In-place Fisher-Yates shuffle.
// O(n) assuming an O(1) PRNG.
func ShuffleFisherYates(ns []int) {
	for i := len(ns) - 1; i >= 1; i-- {
		j := rand.Intn(i + 1)
		ns[i], ns[j] = ns[j], ns[i]
	}
}

// In-place shuffle using a sort based on a random number.
// O(n log n) assuming an O(1) PRNG.
func ShuffleRandomSort(ns []int) {
	sort.Sort(RandomOrderSlice(ns))
}
