package seq

import (
	"iter"
)

// Range generates a sequence of integers from 0 to `size - 1`.
//
// This function creates a sequence that yields integers starting from 0 up to, but not including, `size`.
//
// Parameters:
//
//	size (int): The upper limit (exclusive) of the sequence.
//
// Returns:
//
//	iter.Seq[int]: A sequence of integers from 0 to `size - 1`.
func Range(size int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for num := range size {
			if !yield(num) {
				return
			}
		}
	}
}
