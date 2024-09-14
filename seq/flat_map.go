package seq

import "iter"

// FlatMap function flattens a sequence of slices 'input' into a single sequence of elements of type E.
// It concatenates the elements of each slice in the sequence into a single sequence.
// This function works with slices of any type E.
//
// Parameters:
//
//	input: The sequence of slices to flatten, of type iter.Seq[~[]E].
//
// Returns:
//
//	A new sequence that implements the iter.Seq[E] interface, containing all the elements from the flattened slices.
func FlatMap[S ~[]E, E any](input iter.Seq[S]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for slice := range input {
			for _, item := range slice {
				if !yield(item) {
					return
				}
			}
		}
	}
}
