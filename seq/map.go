package seq

import "iter"

// Map applies a transformation function to each element in the input sequence and returns a new sequence.
//
// This function takes a sequence `input` and a `mapper` function. It applies the `mapper` function
// to each element of the sequence and returns a new sequence with the transformed elements.
//
// Parameters:
//
//	input (iter.Seq[I]): The input sequence of elements.
//	mapper (func(I) O): The transformation function to apply to each element.
//
// Returns:
//
//	iter.Seq[O]: A new sequence with transformed elements.
func Map[I, O any](input iter.Seq[I], mapper func(I) O) iter.Seq[O] {
	return func(yield func(O) bool) {
		for item := range input {
			if !yield(mapper(item)) {
				return
			}
		}
	}
}
