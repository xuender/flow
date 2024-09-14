package seq

import "iter"

// Filter function creates a new sequence that contains only the elements from the given sequence 'input'
// that satisfy the condition defined by the 'predicate' function.
// This function works with elements of any type E.
//
// Parameters:
//
//	input: The sequence to filter, of type iter.Seq[E].
//
// predicate: A function that determines if an element of type E should be included.
//
// Returns:
//
//	A new sequence that implements the iter.Seq[E] interface, containing only the elements that satisfy the predicate.
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
