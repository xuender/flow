package seq

import "iter"

// Emit consumes the input sequence by calling its iterator with a function that always returns true.
//
// This function effectively drains the sequence `input` by iterating over all its elements.
//
// Parameters:
//
//	input (iter.Seq[E]): The input sequence of elements to be consumed.
func Emit[E any](input iter.Seq[E]) {
	input(func(E) bool {
		return true
	})
}
