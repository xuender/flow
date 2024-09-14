package seq

import "iter"

// Limit creates a new sequence containing at most 'limit' elements from 'input'.
// If 'input' has fewer elements, all are included.
// Works with any element type E.
//
// Parameters:
//
//	input: The sequence to limit, of type iter.Seq[E].
//	limit: The maximum number of elements to include.
//
// Returns:
//
//	A new sequence that implements the iter.Seq[E] interface, containing up to 'limit' elements.
func Limit[E any](input iter.Seq[E], limit int) iter.Seq[E] {
	return func(yield func(E) bool) {
		idx := 0

		for item := range input {
			if idx >= limit {
				return
			}

			if !yield(item) {
				return
			}

			idx++
		}
	}
}
