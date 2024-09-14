package seq

import (
	"iter"
)

// Count function is used to calculate the number of elements in a given sequence.
// It accepts a generic sequence 'input' as input, which can handle elements of any type E.
// The function calculates the total number of elements by iterating through the entire sequence.
//
// Parameters:
//
//	input: A generic sequence containing elements of type E.
//
// Returns:
//
//	int: The total number of elements in the sequence.
func Count[E any](input iter.Seq[E]) int {
	count := 0
	for range input {
		count++
	}

	return count
}
