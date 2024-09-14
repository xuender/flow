package seq

import "iter"

// FlatMap flattens a sequence of slices into a single sequence of elements.
//
// This function takes a sequence of slices `input` and returns a new sequence that contains
// all the elements of the slices concatenated together.
//
// Args:
//
//	input (iter.Seq[S]): The input sequence of slices.
//
// Returns:
//
//	iter.Seq[E]: A new sequence containing all the elements of the input slices.
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
