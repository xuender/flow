package seq

import "iter"

// Limit returns a new sequence containing at most `limit` elements from the input sequence.
//
// This function iterates over the sequence `input` and includes only the first `limit` elements
// in the new sequence. If the input sequence has fewer than `limit` elements, the new sequence
// will contain all elements.
func Limit[V any](input iter.Seq[V], limit int) iter.Seq[V] {
	return func(yield func(V) bool) {
		idx := 0

		for item := range input {
			if idx >= limit {
				return
			}

			if !yield(item) {
				return
			}

			idx++
		}
	}
}

// Limit2 creates a new sequence that only includes the first N elements from the input sequence.
//
// Iterates through the input sequence and yields elements until the limit is reached.
// Stops yielding elements once the limit is exceeded.
func Limit2[K, V any](input iter.Seq2[K, V], limit int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		idx := 0

		for key, val := range input {
			if idx >= limit {
				return
			}

			if !yield(key, val) {
				return
			}

			idx++
		}
	}
}
