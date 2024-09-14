package seq

import "iter"

// Each applies a function to each element in the input sequence.
// This function iterates over the sequence `input` and applies the `yield` function to each element.
// Iteration stops if `yield` returns false.
//
// Args:
//
//	input (iter.Seq[V]): The input sequence of elements.
//	yield (func(V) bool): The function to apply to each element.
func Each[V any](input iter.Seq[V], yield func(V) bool) {
	for item := range input {
		if !yield(item) {
			return
		}
	}
}

// Each2 iterates over a sequence of (key, value) pairs and applies a callback function.
// It processes each (key, value) pair in the sequence.

// Args:
//
//	input (iter.Seq2[K, V]): The input sequence of (key, value) pairs.
//	yield (func(K, V) bool): The callback function to apply for each pair.
func Each2[K, V any](input iter.Seq2[K, V], yield func(K, V) bool) {
	for key, val := range input {
		if !yield(key, val) {
			return
		}
	}
}
