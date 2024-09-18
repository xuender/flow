package seq

import "iter"

// Emit consumes the input sequence by calling its iterator with a function that always returns true.
//
// It effectively drains the sequence `input` by iterating over all its elements.
//
// Play: https://go.dev/play/p/YmFQ4CMN5Q_u
func Emit[V any](input iter.Seq[V]) {
	input(func(V) bool {
		return true
	})
}

// Emit2 iterates over a sequence of (key, value) pairs and discards each element.
//
// It processes each (key, value) pair without returning any value.
//
// Play: https://go.dev/play/p/Sc41jwY3k5__W
func Emit2[K, V any](input iter.Seq2[K, V]) {
	input(func(K, V) bool {
		return true
	})
}
