package seq

import (
	"iter"
)

// Keys extracts keys from a sequence of (key, value) pairs.
//
// It returns a new sequence containing only the keys.
//
// Play: https://go.dev/play/p/IlQDB_yWQ8Z
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
//
// Play: https://go.dev/play/p/X0ZD9Y5qxB8
func Values[K, V any](input iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, val := range input {
			if !yield(val) {
				return
			}
		}
	}
}
