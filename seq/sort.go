package seq

import (
	"cmp"
	"iter"
	"slices"
)

// Sort sorts a sequence of elements that implement the cmp.Ordered interface.
//
// It takes a sequence `input` of orderable elements and returns a sorted sequence.
// The returned value is an iterator function that, when called with a `yield` function,
// iterates over the sorted elements.
//
// Parameters:
//
//	input (iter.Seq[E]): A sequence of elements implementing cmp.Ordered to be sorted.
//
// Returns:
//
//	iter.Seq[E]: An iterator function that yields sorted elements.
func Sort[E cmp.Ordered](input iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		slice := slices.Collect(input)
		slices.Sort(slice)

		for _, item := range slice {
			if !yield(item) {
				return
			}
		}
	}
}

// SortFunc creates a new sequence with elements from 'input' sorted using the custom comparison function 'cmp'.
// Works with any element type E.
//
// Parameters:
//
//	input: The sequence to sort, of type iter.Seq[E].
//	cmp: A comparison function that defines the order of the elements.
//
// Returns:
//
//	A new sequence that implements the iter.Seq[E] interface, containing sorted elements.
func SortFunc[E any](input iter.Seq[E], cmp func(item1, item2 E) int) iter.Seq[E] {
	return func(yield func(E) bool) {
		slice := slices.Collect(input)
		slices.SortFunc(slice, cmp)

		for _, item := range slice {
			if !yield(item) {
				return
			}
		}
	}
}
