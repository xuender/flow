package seq

import "iter"

// Peek applies a given action to each element in the input sequence and returns the original sequence.
//
// This function takes a sequence `input` and an `action` function. It applies the `action` to each element
// of the sequence and then returns the original sequence as an iterator.
//
// Args:
//
//	input iter.Seq[V]: The input sequence of elements.
//	action func(V): A function to apply to each element.
//
// Returns:
//
//	iter.Seq[V]: An iterator function yielding the original elements.
func Peek[V any](input iter.Seq[V], action func(V)) iter.Seq[V] {
	return func(yield func(V) bool) {
		for item := range input {
			action(item)

			if !yield(item) {
				return
			}
		}
	}
}

// Peek2 applies an action to each (key, value) pair in the sequence.
//
// It returns a new sequence with the same pairs.
//
// Args:
//
//	input iter.Seq2[K, V]: The input sequence of (key, value) pairs.
//	action func(K, V): The action to apply to each pair.
//
// Returns:
//
//	iter.Seq2[K, V]: A new sequence with the same (key, value) pairs.
func Peek2[K, V any](input iter.Seq2[K, V], action func(K, V)) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for key, val := range input {
			action(key, val)

			if !yield(key, val) {
				return
			}
		}
	}
}
