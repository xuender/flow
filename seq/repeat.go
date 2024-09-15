package seq

import (
	"iter"
)

// Repeat creates a new sequence by repeating each element of the input sequence a specified number of times.
//
// This function returns a new sequence where each element of the input sequence is repeated `count` times.
//
// Args:
//
//	input iter.Seq[V]: The input sequence.
//	count int: The number of times to repeat each element.
//
// Returns:
//
//	iter.Seq[V]: A new sequence with repeated elements.
func Repeat[V any](input iter.Seq[V], count int) iter.Seq[V] {
	return func(yield func(V) bool) {
		for item := range input {
			for range count {
				if !yield(item) {
					return
				}
			}
		}
	}
}

// Repeat2 creates a new sequence by repeating each element of the input sequence a specified number of times.
//
// This function returns a new sequence where each element of the input sequence is repeated `count` times.
//
// Args:
//
//	input iter.Seq2[K, V]: The input sequence.
//	count int: The number of times to repeat each element.
//
// Returns:
//
//	iter.Seq2[K, V]: A new sequence with repeated elements.
func Repeat2[K, V any](input iter.Seq2[K, V], count int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for key, val := range input {
			for range count {
				if !yield(key, val) {
					return
				}
			}
		}
	}
}