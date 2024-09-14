package seq

import (
	"iter"
	"slices"
)

// Reverse reverses the elements of the input sequence and returns them as an iterator.
//
// This function takes a sequence `input` and returns a new iterator that yields the elements
// of the original sequence in reverse order.
//
// Parameters:
//
//	input (iter.Seq[E]): The input sequence of elements to be reversed.
//
// Returns:
//
//	iter.Seq[E]: An iterator function yielding elements in reverse order.
func Reverse[E any](input iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		slice := slices.Collect(input)
		slices.Reverse(slice)

		for _, item := range slice {
			if !yield(item) {
				return
			}
		}
	}
}
