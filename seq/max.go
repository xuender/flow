package seq

import (
	"cmp"
	"iter"
)

// Max returns the maximum element in the input sequence or false if the sequence is empty.
//
// If the sequence is empty, it returns the zero value of type `E` and false.
//
// Play: https://go.dev/play/p/1HUDnJ2ofC-
func Max[V cmp.Ordered](input iter.Seq[V]) (V, bool) {
	var (
		maxVal V
		has    bool
	)

	for item := range input {
		if !has || item > maxVal {
			maxVal = item
			has = true
		}
	}

	return maxVal, has
}

// Max2 finds the maximum (key, value) pair in the sequence.
//
// It returns the maximum key, value, and a boolean indicating if a maximum was found.
//
// Play: https://go.dev/play/p/eAbCubpT5mS
func Max2[K cmp.Ordered, V any](input iter.Seq2[K, V]) (K, V, bool) {
	var (
		maxKey K
		maxVal V
		has    bool
	)

	for key, val := range input {
		if !has || key > maxKey {
			maxKey = key
			maxVal = val
			has = true
		}
	}

	return maxKey, maxVal, has
}
