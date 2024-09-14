package seq

import (
	"iter"
	"slices"
)

// Reverse reverses the elements of the input sequence and returns them as an iterator.
//
// This function takes a sequence `input` and returns a new iterator that yields the elements
// of the original sequence in reverse order.
//
// Args:
//
//	input iter.Seq[V]: The input sequence of elements to be reversed.
//
// Returns:
//
//	iter.Seq[V]: An iterator function yielding elements in reverse order.
func Reverse[V any](input iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		slice := slices.Collect(input)
		slices.Reverse(slice)

		for _, item := range slice {
			if !yield(item) {
				return
			}
		}
	}
}

// Reverse2 reverses the order of (key, value) pairs in the sequence.
//
// It returns a new sequence with reversed pairs.
//
// Args:
//
//	input iter.Seq2[K, V]: The input sequence of (key, value) pairs.
//
// Returns:
//
//	iter.Seq2[K, V]: A new sequence with reversed (key, value) pairs.
func Reverse2[K, V any](input iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		items := []Tuple[K, V]{}
		for key, val := range input {
			items = append(items, T(key, val))
		}

		slices.Reverse(items)

		for _, item := range items {
			if !yield(item.A, item.B) {
				return
			}
		}
	}
}
