package seq

import "iter"

// Emit function is used to iterate over a given sequence 'input'.
// It invokes a provided function for each element in the sequence, but in this case,
// the provided function simply returns true, effectively doing nothing with the elements.
// This function works with elements of any type E.
//
// Parameters:
//
//	input: The sequence to iterate over, of type iter.Seq[E].
//
// Returns:
//
//	No return value.
func Emit[E any](input iter.Seq[E]) {
	input(func(E) bool {
		return true
	})
}
