package seq

import (
	"iter"
	"math/rand/v2"
	"slices"
)

// Shuffle returns a new sequence with the elements of the input sequence in a random order.
//
// It shuffles the input and yields each element in the shuffled order.
//
// Play: https://go.dev/play/p/dVl8GzLyL6l
func Shuffle[V any](input iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		slice := slices.Collect(input)

		rand.Shuffle(len(slice), func(i, j int) {
			slice[i], slice[j] = slice[j], slice[i]
		})

		for _, item := range slice {
			if !yield(item) {
				return
			}
		}
	}
}

// Shuffle2 returns a new sequence with the pairs of the input sequence in a random order.
//
// It shuffles the input pairs and yields each pair in the shuffled order.
//
// Play: https://go.dev/play/p/vvtLAWPy2Nf
func Shuffle2[K, V any](input iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		slice := Tuples(input)

		rand.Shuffle(len(slice), func(i, j int) {
			slice[i], slice[j] = slice[j], slice[i]
		})

		for _, item := range slice {
			if !yield(item.K, item.V) {
				return
			}
		}
	}
}
