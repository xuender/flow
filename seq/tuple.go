package seq

import "iter"

// Tuple represents a pair of values.
//
// It contains two values of different types.
type Tuple[V1, V2 any] struct {
	A V1
	B V2
}

// T creates a Tuple from two values.
//
// It returns a Tuple containing the values.
func T[V1, V2 any](val1 V1, val2 V2) Tuple[V1, V2] {
	return Tuple[V1, V2]{val1, val2}
}

// Tuples converts a sequence of (key, value) pairs into a slice of Tuple[K, V].
//
// It returns the slice containing the tuples.
func Tuples[K, V any](input iter.Seq2[K, V]) []Tuple[K, V] {
	ret := []Tuple[K, V]{}

	for key, val := range input {
		ret = append(ret, T(key, val))
	}

	return ret
}
