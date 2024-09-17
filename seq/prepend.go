package seq

import "iter"

// Prepend adds items to the beginning of the sequence.
//
// It function is useful for prepending items to a sequence in a functional manner.
func Prepend[V any](input iter.Seq[V], items ...V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, item := range items {
			if !yield(item) {
				return
			}
		}

		for item := range input {
			if !yield(item) {
				return
			}
		}
	}
}

// Prepend2 adds tuples to the beginning of the sequence.
//
// It function is useful for prepending tuples to a sequence in a functional manner.
func Prepend2[K, V any](input iter.Seq2[K, V], items ...Tuple[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, item := range items {
			if !yield(item.K, item.V) {
				return
			}
		}

		for key, val := range input {
			if !yield(key, val) {
				return
			}
		}
	}
}
