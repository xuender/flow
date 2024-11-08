package seq

import (
	"iter"
	"slices"
)

// Seq creates and returns a sequence iterator containing the given elements.
//
// It returns an iter.Seq[V] that can be used to iterate over the provided elements.
func Seq[V any](values ...V) iter.Seq[V] {
	return slices.Values(values)
}

// Seq2 creates and returns a sequence iterator for key-value pairs.
//
// It returns an iter.Seq2[K, V] that can be used to iterate over the provided key-value pairs.
func Seq2[K, V any](values ...Tuple[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, val := range values {
			if !yield(val.K, val.V) {
				return
			}
		}
	}
}
