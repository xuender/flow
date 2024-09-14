package seq

import (
	"iter"
)

// Slice converts the input sequence into a slice of elements with a specified capacity.
//
// This function iterates over the sequence `input` and collects the elements into a slice.
// The slice has a capacity of `size`.
//
// Parameters:
//
//	input (iter.Seq[E]): The input sequence of elements.
//	size (int): The capacity of the resulting slice.
//
// Returns:
//
//	[]E: A slice containing the elements of the sequence.
func Slice[E any](input iter.Seq[E], size int) []E {
	ret := make([]E, 0, size)

	for item := range input {
		ret = append(ret, item)
	}

	return ret
}
