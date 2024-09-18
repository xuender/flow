package seq

import (
	"cmp"
	"iter"
)

// Sum calculates the sum of all elements in the input sequence.
//
// It iterates over the sequence `input` and calculates the sum of all elements.
//
// Play: https://go.dev/play/p/EL_lsMnNq7u
func Sum[V cmp.Ordered](input iter.Seq[V]) V {
	var sum V
	for item := range input {
		sum += item
	}

	return sum
}
