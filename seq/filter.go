package seq

import "iter"

// Filter returns a new sequence containing only the elements that satisfy the given predicate.
//
// This function iterates over the sequence `input` and applies the `predicate` function to each element.
// Only elements for which the predicate returns true are included in the new sequence.
//
// Args:
//
//	input (iter.Seq[E]): The input sequence of elements.
//	predicate (func(E) bool): The predicate function to filter elements.
//
// Returns:
//
//	iter.Seq[E]: A new sequence containing filtered elements.
func Filter[E any](input iter.Seq[E], predicate func(E) bool) iter.Seq[E] {
	return func(yield func(E) bool) {
		for item := range input {
			if predicate(item) {
				if !yield(item) {
					return
				}
			}
		}
	}
}
