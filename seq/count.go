package seq

import (
	"iter"
)

// Count returns the number of elements in the input sequence.
//
// This function iterates over the sequence `input` and counts the number of elements.
//
// Args:
//
//	input (iter.Seq[V]): The input sequence of elements.
//
// Returns:
//
//	int: The number of elements in the sequence.
func Count[V any](input iter.Seq[V]) int {
	count := 0
	for range input {
		count++
	}

	return count
}

// Count2 returns the number of elements in the input sequence.
//
// This function iterates over the sequence `input` and counts the number of elements.
//
// Args:
//
//	input (iter.Seq2[K, V]): The input sequence of elements.
//
// Returns:
//
//	int: The number of elements in the sequence.
func Count2[K, V any](input iter.Seq2[K, V]) int {
	count := 0
	for range input {
		count++
	}

	return count
}
