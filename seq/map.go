package seq

import "iter"

// Map function transforms each element of the given sequence 'input' using the provided 'mapper' function.
// It creates a new sequence that contains the transformed elements.
// This function works with elements of any input type I and output type O.
//
// Parameters:
//
//	input: The sequence to transform, of type iter.Seq[I].
//	mapper: A function that takes an element of type I and returns an element of type O.
//
// Returns:
//
//	A new sequence that implements the iter.Seq[O] interface, containing the transformed elements.
func Map[I, O any](input iter.Seq[I], mapper func(I) O) iter.Seq[O] {
	return func(yield func(O) bool) {
		for item := range input {
			if !yield(mapper(item)) {
				return
			}
		}
	}
}
