package seq

import "iter"

// Skip skips the first `offset` elements of the input sequence and returns the remaining elements as an iterator.
//
// This function takes a sequence `input` and an integer `offset`, then creates a new iterator that
// skips the first `offset` elements from the original sequence.
//
// Args:
//
//	input iter.Seq[V]: The input sequence of elements.
//	offset int: The number of elements to skip from the beginning of the sequence.
//
// Returns:
//
//	iter.Seq[V]: An iterator function yielding elements after the specified offset.
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
// Args:
//
//	input iter.Seq2[K, V]: The input sequence of (key, value) pairs.
//	offset int: The number of pairs to skip.
//
// Returns:
//
//	iter.Seq2[K, V]: A new sequence with skipped pairs.
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
