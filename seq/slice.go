package seq

import (
	"iter"
)

// Slice function collects the elements of the given sequence 'input' into a slice of size 'size'.
// It returns a slice containing the elements of the sequence.
// This function works with elements of any type E.
//
// Parameters:
//
//	input: The sequence to collect elements from, of type iter.Seq[E].
//	size: The capacity hint for the slice to be created.
//
// Returns:
//
//	A slice of type E containing the elements of the sequence.
func Slice[E any](input iter.Seq[E], size int) []E {
	ret := make([]E, 0, size)

	for item := range input {
		ret = append(ret, item)
	}

	return ret
}
