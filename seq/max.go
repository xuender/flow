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

// Finds the maximum value in a sequence using a custom comparison function.
//
// Returns the maximum value and a boolean indicating if a maximum was found.
func MaxFunc[V any](input iter.Seq[V], cmp func(V, V) int) (V, bool) {
	var (
		maxVal V
		has    bool
	)

	for item := range input {
		if !has || cmp(item, maxVal) > 0 {
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

// MaxFunc2 finds the maximum key-value pair in a sequence using a provided comparison function.
//
// Returns the maximum key, value, and a boolean indicating if a maximum was found.
func MaxFunc2[K, V any](input iter.Seq2[K, V], cmp func(Tuple[K, V], Tuple[K, V]) int) (K, V, bool) {
	var (
		maxKey K
		maxVal V
		has    bool
	)

	for key, val := range input {
		if !has || cmp(T(key, val), T(maxKey, maxVal)) > 0 {
			maxKey = key
			maxVal = val
			has = true
		}
	}

	return maxKey, maxVal, has
}
