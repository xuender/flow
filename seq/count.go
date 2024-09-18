package seq

import (
	"iter"
)

// Count returns the number of elements in the input sequence.
//
// It iterates over the sequence `input` and counts the number of elements.
//
// Play: https://go.dev/play/p/qtsl1eaa1Sk
func Count[V any](input iter.Seq[V]) int {
	count := 0
	for range input {
		count++
	}

	return count
}

// Count2 returns the number of elements in the input sequence.
//
// It iterates over the sequence `input` and counts the number of elements.
//
// Play: https://go.dev/play/p/ntU3HJ1wfAX
func Count2[K, V any](input iter.Seq2[K, V]) int {
	count := 0
	for range input {
		count++
	}

	return count
}
