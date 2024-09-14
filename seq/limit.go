package seq

import "iter"

// Limit returns a new sequence containing at most `limit` elements from the input sequence.
//
// This function iterates over the sequence `input` and includes only the first `limit` elements
// in the new sequence. If the input sequence has fewer than `limit` elements, the new sequence
// will contain all elements.
//
// Args:
//
//	input (iter.Seq[E]): The input sequence of elements.
//	limit (int): The maximum number of elements to include.
//
// Returns:
//
//	iter.Seq[E]: A new sequence with at most `limit` elements.
func Limit[E any](input iter.Seq[E], limit int) iter.Seq[E] {
	return func(yield func(E) bool) {
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
