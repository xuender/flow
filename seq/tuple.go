package seq

import "iter"

// Tuple represents a pair of values.
//
// It contains two values of different types.
type Tuple[K, V any] struct {
	K K
	V V
}

// T creates a Tuple from two values.
//
// It returns a Tuple containing the values.
//
// Play: https://go.dev/play/p/yGCbA5v9OyG
func T[K, V any](key K, val V) Tuple[K, V] {
	return Tuple[K, V]{key, val}
}

// Tuples converts a sequence of (key, value) pairs into a slice of Tuple[K, V].
//
// It returns the slice containing the tuples.
//
// Play: https://go.dev/play/p/t9QhSnSONUB
func Tuples[K, V any](input iter.Seq2[K, V]) []Tuple[K, V] {
	ret := []Tuple[K, V]{}

	for key, val := range input {
		ret = append(ret, T(key, val))
	}

	return ret
}
