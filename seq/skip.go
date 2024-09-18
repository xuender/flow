package seq

import "iter"

// Skip skips the first `offset` elements of the input sequence and returns the remaining elements as an iterator.
//
// It takes a sequence `input` and an integer `offset`, then creates a new iterator that
// skips the first `offset` elements from the original sequence.
//
// Play: https://go.dev/play/p/Oxf000GM6Sx
func Skip[V any](input iter.Seq[V], offset int) iter.Seq[V] {
	return func(yield func(V) bool) {
		idx := 0

		for item := range input {
			idx++
			if idx <= offset {
				continue
			}

			if !yield(item) {
				return
			}
		}
	}
}

// Skip2 skips the first 'offset' (key, value) pairs in the sequence.
//
// It returns a new sequence with the remaining pairs.
//
// Play: https://go.dev/play/p/wipik9SWE6Y
func Skip2[K, V any](input iter.Seq2[K, V], offset int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		idx := 0

		for key, val := range input {
			idx++
			if idx <= offset {
				continue
			}

			if !yield(key, val) {
				return
			}
		}
	}
}
