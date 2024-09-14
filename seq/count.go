package seq

import (
	"iter"
)

// Count returns the number of elements in the input sequence.
//
// This function iterates over the sequence `input` and counts the number of elements.
//
// Parameters:
//
//	input (iter.Seq[E]): The input sequence of elements.
//
// Returns:
//
//	int: The number of elements in the sequence.
func Count[E any](input iter.Seq[E]) int {
	count := 0
	for range input {
		count++
	}

	return count
}
