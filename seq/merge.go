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
			Each(seq, yield)
		}
	}
}
