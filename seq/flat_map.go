package seq

import "iter"

// FlatMap flattens a sequence of slices into a single sequence of elements.
//
// It takes a sequence of slices `input` and returns a new sequence that contains
// all the elements of the slices concatenated together.
func FlatMap[S ~[]V, V any](input iter.Seq[S]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for slice := range input {
			for _, item := range slice {
				if !yield(item) {
					return
				}
			}
		}
	}
}
