package seq

import "iter"

// Filter returns a new sequence containing only the elements that satisfy the given predicate.
//
// It iterates over the sequence `input` and applies the `predicate` function to each element.
// Only elements for which the predicate returns true are included in the new sequence.
func Filter[V any](input iter.Seq[V], predicate func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for item := range input {
			if predicate(item) {
				if !yield(item) {
					return
				}
			}
		}
	}
}

// Filter2 filters a sequence of (key, value) pairs based on a predicate.
//
// It returns a new sequence containing only the pairs that satisfy the predicate.
func Filter2[K, V any](input iter.Seq2[K, V], predicate func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for key, val := range input {
			if predicate(key, val) {
				if !yield(key, val) {
					return
				}
			}
		}
	}
}
