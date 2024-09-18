package seq

import "iter"

// Each applies a function to each element in the input sequence.
//
// It iterates over the sequence `input` and applies the `yield` function to each element.
// Iteration stops if `yield` returns false.
//
// Play: https://go.dev/play/p/Al9R-6G1GJI
func Each[V any](input iter.Seq[V], yield func(V) bool) {
	for item := range input {
		if !yield(item) {
			return
		}
	}
}

// Each2 iterates over a sequence of (key, value) pairs and applies a callback function.
//
// It processes each (key, value) pair in the sequence.
//
// Play: https://go.dev/play/p/z_ZjLirSPnO
func Each2[K, V any](input iter.Seq2[K, V], yield func(K, V) bool) {
	for key, val := range input {
		if !yield(key, val) {
			return
		}
	}
}
