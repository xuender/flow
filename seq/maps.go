package seq

import (
	"iter"
)

// Keys extracts keys from a sequence of (key, value) pairs.
//
// It returns a new sequence containing only the keys.
func Keys[K, V any](input iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for key := range input {
			if !yield(key) {
				return
			}
		}
	}
}

// Values extracts values from a sequence of (key, value) pairs.
//
// It returns a new sequence containing only the values.
func Values[K, V any](input iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, val := range input {
			if !yield(val) {
				return
			}
		}
	}
}
