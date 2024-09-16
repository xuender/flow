package seq

import (
	"cmp"
	"iter"
)

// Sum calculates the sum of all elements in the input sequence.
//
// This function iterates over the sequence `input` and calculates the sum of all elements.
func Sum[E cmp.Ordered](input iter.Seq[E]) E {
	var sum E
	for item := range input {
		sum += item
	}

	return sum
}
