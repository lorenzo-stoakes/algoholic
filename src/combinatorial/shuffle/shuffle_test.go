package algoholic

import "testing"

const BENCHMARK_LENGTH = 10

func TestFisherYatesShufflesEvenly(t *testing.T) {
	// Set an arbitrary maximum error from expectation - We expect each number to occur within
	// 1% of the probability in a list of length 10 after 1e6 iterations.
	// TODO: Determine maximum error mathematically.
	checkShufflesEvenly(t, ShuffleFisherYates, 10, 1e6, 0.01)
}

func TestRandomOrderShuffleShufflesEvenly(t *testing.T) {
	// Set an arbitrary maximum error from expectation - We expect each number to occur within
	// 1% of the probability in a list of length 10 after 1e6 iterations.
	// TODO: Determine maximum error mathematically.
	checkShufflesEvenly(t, ShuffleRandomSort, 10, 1e6, 0.01)
}

func BenchmarkFisherYatesShuffle(b *testing.B) {
	benchmarkShuffle(b, BENCHMARK_LENGTH, ShuffleFisherYates)
}

func BenchmarkRandomSortShuffle(b *testing.B) {
	benchmarkShuffle(b, BENCHMARK_LENGTH, ShuffleRandomSort)
}
