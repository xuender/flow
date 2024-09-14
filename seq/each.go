package seq

import "iter"

// Each applies a function to each element in the input sequence.
// This function iterates over the sequence `input` and applies the `yield` function to each element.
// Iteration stops if `yield` returns false.
//
// Args:
//
//	input (iter.Seq[E]): The input sequence of elements.
//	yield (func(E) bool): The function to apply to each element.
func Each[E any](input iter.Seq[E], yield func(E) bool) {
	for item := range input {
		if !yield(item) {
			return
		}
	}
}
