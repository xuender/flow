package seq

import (
	"cmp"
	"iter"
	"maps"
	"slices"
)

// Sort sorts a sequence of elements that implement the cmp.Ordered interface.
//
// It takes a sequence `input` of orderable elements and returns a sorted sequence.
// The returned value is an iterator function that, when called with a `yield` function,
// iterates over the sorted elements.
//
// Args:
//
//	input (iter.Seq[V]): A sequence of elements implementing cmp.Ordered to be sorted.
//
// Returns:
//
//	iter.Seq[V]: An iterator function that yields sorted elements.
func Sort[V cmp.Ordered](input iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		slice := slices.Collect(input)
		slices.Sort(slice)

		for _, item := range slice {
			if !yield(item) {
				return
			}
		}
	}
}

// Sort2 sorts the input sequence by keys and returns a sorted sequence.
//
// It takes a sequence of (key, value) pairs where keys are ordered.
//
// Args:
//
//	input iter.Seq2[K, V]: The input sequence of (key, value) pairs.
//
// Returns:
//
//	iter.Seq2[K, V]: A new sequence yielding sorted (key, value) pairs.
func Sort2[K cmp.Ordered, V any](input iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		items := maps.Collect(input)

		for key := range Sort(maps.Keys(items)) {
			if !yield(key, items[key]) {
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
// Args:
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
