package seq

import (
	"iter"
)

// Distinct returns a new sequence containing only the distinct elements from the input sequence.
//
// This function iterates over the sequence `input` and filters out duplicate elements.
// Only the first occurrence of each element is included in the new sequence.
//
// Parameters:
//
//	input (iter.Seq[E]): The input sequence of comparable elements.
//
// Returns:
//
//	iter.Seq[E]: A new sequence with distinct elements.
func Distinct[E comparable](input iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		none := struct{}{}
		set := map[E]struct{}{}

		for item := range input {
			if _, has := set[item]; has {
				continue
			}

			set[item] = none

			if !yield(item) {
				return
			}
		}
	}
}
