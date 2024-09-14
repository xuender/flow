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

// SortFunc sorts the elements of the input sequence using a custom comparison function.
//
// This function collects the elements of `input` into a slice, sorts them using the `cmp` function,
// and returns a new sorted sequence.
//
// Parameters:
//
//	input (iter.Seq[E]): The input sequence of elements.
//	cmp (func(item1, item2 E) int): The comparison function for sorting.
//
// Returns:
//
//	iter.Seq[E]: A new sequence with sorted elements.
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
