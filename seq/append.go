package seq

import "iter"

// Append appends multiple elements to the end of the input sequence.

// It takes an input sequence and additional elements to append.

// Args:
//
//	input iter.Seq[V]: The input sequence.
//	items ...V: Elements to append to the sequence.
//
// Returns:
//
//	iter.Seq[V]: A new sequence with the appended elements.
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
