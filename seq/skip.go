package seq

import "iter"

// Skip skips the first `offset` elements of the input sequence and returns the remaining elements as an iterator.
//
// This function takes a sequence `input` and an integer `offset`, then creates a new iterator that
// skips the first `offset` elements from the original sequence.
//
// Args:
//
//	input (iter.Seq[E]): The input sequence of elements.
//	offset (int): The number of elements to skip from the beginning of the sequence.
//
// Returns:
//
//	iter.Seq[E]: An iterator function yielding elements after the specified offset.
func Skip[E any](input iter.Seq[E], offset int) iter.Seq[E] {
	return func(yield func(E) bool) {
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
