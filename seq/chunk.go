package seq

import (
	"iter"
	"slices"
)

func Chunk[V any](input iter.Seq[V], num int) []iter.Seq[V] {
	seqs := slices.Chunk(slices.Collect(input), num)
	ret := make([]iter.Seq[V], 0, num)

	for items := range seqs {
		ret = append(ret, slices.Values(items))
	}

	return ret
}

func Chunk2[K, V any](input iter.Seq2[K, V], num int) []iter.Seq2[K, V] {
	seqs := slices.Chunk(Tuples(input), num)
	ret := make([]iter.Seq2[K, V], 0, num)

	for items := range seqs {
		ret = append(ret, func(yield func(K, V) bool) {
			for _, item := range items {
				if !yield(item.A, item.B) {
					return
				}
			}
		})
	}

	return ret
}
