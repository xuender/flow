package seq

import "iter"

// Concat combines multiple sequence iterators into a single sequence iterator.
//
// It returns an iter.Seq[V] that iterates over all elements from the provided sequences.
func Concat[V any](seqs ...iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, seq := range seqs {
			for val := range seq {
				if !yield(val) {
					return
				}
			}
		}
	}
}

// Concat2 combines multiple key-value sequence iterators into a single sequence iterator.
//
// It returns an iter.Seq2[K, V] that iterates over all key-value pairs from the provided sequences.
func Concat2[K, V any](seqs ...iter.Seq2[K, V]) iter.Seq2[K, V] {
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
