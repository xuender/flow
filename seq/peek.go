package seq

import "iter"

// Peek applies a given action to each element in the input sequence and returns the original sequence.
//
// This function takes a sequence `input` and an `action` function. It applies the `action` to each element
// of the sequence and then returns the original sequence as an iterator.
//
// Args:
//
//	input (iter.Seq[E]): The input sequence of elements.
//	action (func(E)): A function to apply to each element.
//
// Returns:
//
//	iter.Seq[E]: An iterator function yielding the original elements.
func Peek[E any](input iter.Seq[E], action func(E)) iter.Seq[E] {
	return func(yield func(E) bool) {
		for item := range input {
			action(item)

			if !yield(item) {
				return
			}
		}
	}
}
