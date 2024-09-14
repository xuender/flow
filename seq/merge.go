package seq

import "iter"

// Merge combines multiple sequences into a single sequence.
//
// It takes multiple sequences and merges them sequentially.
//
// Args:
//
//	seqs ...iter.Seq[V]: Multiple sequences to merge.
//
// Returns:
//
//	iter.Seq[V]: A new sequence containing all elements from the input sequences.
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
// Args:
//
//	seqs ...iter.Seq2[K, V]: Multiple sequences of (key, value) pairs to merge.
//
// Returns:
//
//	iter.Seq2[K, V]: A new sequence containing all (key, value) pairs.
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
