package seq

import "iter"

// Append appends multiple elements to the end of the input sequence.
//
// It takes an input sequence and additional elements to append.
func Append[V any](input iter.Seq[V], items ...V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for item := range input {
			if !yield(item) {
				return
			}
		}

		for _, item := range items {
			if !yield(item) {
				return
			}
		}
	}
}

// Append2 appends additional (key, value) pairs to the input sequence.
//
// It returns a new sequence with the appended pairs.
func Append2[K, V any](input iter.Seq2[K, V], items ...Tuple[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for key, val := range input {
			if !yield(key, val) {
				return
			}
		}

		for _, item := range items {
			if !yield(item.K, item.V) {
				return
			}
		}
	}
}
