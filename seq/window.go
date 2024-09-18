package seq

import (
	"iter"
)

// Window creates a sequence of sliding windows over the input sequence.
//
// Each window contains `size` consecutive elements from the input.
//
// Play: https://go.dev/play/p/tpADrJdzuSC
func Window[V any](input iter.Seq[V], size int) iter.Seq[[]V] {
	return func(yield func([]V) bool) {
		cache := []V{}

		for item := range input {
			cache = append(cache, item)

			if len(cache) > size {
				cache = cache[1:]
			}

			if !yield(cache) {
				return
			}
		}
	}
}

// Window2 creates a sequence of sliding windows over pairs in the input sequence.
//
// Each window contains `size` consecutive pairs from the input.
func Window2[K, V any](input iter.Seq2[K, V], size int) iter.Seq[[]Tuple[K, V]] {
	return func(yield func([]Tuple[K, V]) bool) {
		cache := []Tuple[K, V]{}

		for key, val := range input {
			cache = append(cache, T(key, val))

			if len(cache) > size {
				cache = cache[1:]
			}

			if !yield(cache) {
				return
			}
		}
	}
}
