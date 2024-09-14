package seq

import (
	"iter"
)

// Range function generates a sequence of integers from 0 (inclusive) to 'size' (exclusive).
// It returns a new sequence that implements the iter.Seq[int] interface.
//
// Parameters:
//
//	size: The upper limit of the sequence (non-inclusive).
//
// Returns:
//
//	A new sequence that implements the iter.Seq[int] interface, containing integers from 0 to size-1.
func Range(size int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for num := range size {
			if !yield(num) {
				return
			}
		}
	}
}
