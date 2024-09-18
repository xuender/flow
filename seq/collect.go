package seq

import (
	"iter"
	"maps"
)

// Collect converts the input sequence into a slice of elements with a specified capacity.
//
// It iterates over the sequence `input` and collects the elements into a slice.
// The slice has a capacity of `size`.
func Collect[V any](input iter.Seq[V], size int) []V {
	ret := make([]V, 0, size)

	for item := range input {
		ret = append(ret, item)
	}

	return ret
}

// Collect2 collects (key, value) pairs from the sequence into a map.
//
// It returns a map containing the collected pairs.
func Collect2[K comparable, V any](input iter.Seq2[K, V], size int) map[K]V {
	ret := make(map[K]V, size)

	maps.Insert(ret, input)

	return ret
}
