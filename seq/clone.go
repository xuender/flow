package seq

import (
	"iter"
	"slices"
)

// Clone creates multiple independent iterators from a single input iterator.
//
// This function returns a slice of `Seq[V]` objects, each of which is an independent iterator
// over the elements of the input sequence. The input sequence is only iterated once.
func Clone[V any](input iter.Seq[V], num int) []iter.Seq[V] {
	items := slices.Collect(input)
	ret := make([]iter.Seq[V], num)

	for idx := range num {
		ret[idx] = slices.Values(items)
	}

	return ret
}

// Clone2 creates multiple independent iterators from a single input iterator for key-value pairs.
//
// This function returns a slice of `Seq2[K, V]` objects, each of which is an independent iterator
// over the elements of the input sequence. The input sequence is only iterated once.
func Clone2[K, V any](input iter.Seq2[K, V], num int) []iter.Seq2[K, V] {
	items := Tuples(input)
	ret := make([]iter.Seq2[K, V], num)

	for idx := range num {
		ret[idx] = func(yield func(K, V) bool) {
			for _, item := range items {
				if !yield(item.A, item.B) {
					return
				}
			}
		}
	}

	return ret
}
