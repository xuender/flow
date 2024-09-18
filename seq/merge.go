package seq

import "iter"

// Merge combines multiple sequences into a single sequence.
//
// It takes multiple sequences and merges them sequentially.
//
// Play: https://go.dev/play/p/bOWQQiA4I6A
func Merge[V any](seqs ...iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, seq := range seqs {
			for item := range seq {
				if !yield(item) {
					return
				}
			}
		}
	}
}

// Merge2 combines multiple sequences of (key, value) pairs into a single sequence.
//
// It merges the input sequences sequentially.
//
// Play: https://go.dev/play/p/lfUCSSnrVtd
func Merge2[K, V any](seqs ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, seq := range seqs {
			for key, val := range seq {
				if !yield(key, val) {
					return
				}
			}
		}
	}
}
