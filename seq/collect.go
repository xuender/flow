package seq

import (
	"iter"
	"maps"
)

// Collect converts the input sequence into a slice of elements with a specified capacity.
//
// This function iterates over the sequence `input` and collects the elements into a slice.
// The slice has a capacity of `size`.
//
// Args:
//
//	input iter.Seq[v]: The input sequence of elements.
//	size int: The capacity of the resulting slice.
//
// Returns:
//
//	[]V: A slice containing the elements of the sequence.
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
//
// Args:
//
//	input iter.Seq2[K, V]: The input sequence of (key, value) pairs.
//	size int: The initial capacity of the map.
//
// Returns:
//
//	map[K]V: A map containing the collected (key, value) pairs.
func Collect2[K comparable, V any](input iter.Seq2[K, V], size int) map[K]V {
	ret := make(map[K]V, size)

	maps.Insert(ret, input)

	return ret
}
