package seq

import (
	"iter"
	"slices"
)

// Chunk divides the input sequence into sub-sequences of a specified size.
//
// It returns a slice of sequences, each of the specified size or smaller for the last chunk.
func Chunk[V any](input iter.Seq[V], num int) []iter.Seq[V] {
	seqs := slices.Chunk(slices.Collect(input), num)
	ret := make([]iter.Seq[V], 0, num)

	for items := range seqs {
		ret = append(ret, slices.Values(items))
	}

	return ret
}

// Chunk2 divides the input sequence of key-value pairs into sub-sequences of a specified size.
//
// It returns a slice of sequences, each containing key-value pairs of the specified size or smaller for the last chunk.
func Chunk2[K, V any](input iter.Seq2[K, V], num int) []iter.Seq2[K, V] {
	seqs := slices.Chunk(Tuples(input), num)
	ret := make([]iter.Seq2[K, V], 0, num)

	for items := range seqs {
		ret = append(ret, func(yield func(K, V) bool) {
			for _, item := range items {
				if !yield(item.K, item.V) {
					return
				}
			}
		})
	}

	return ret
}
