package seq

import (
	"cmp"
	"iter"
)

// Min returns the minimum element in the input sequence or false if the sequence is empty.
//
// It iterates over the sequence `input` and finds the minimum element.
// If the sequence is empty, it returns the zero value of type `E` and false.
func Min[V cmp.Ordered](input iter.Seq[V]) (V, bool) {
	var minItem V

	has := false

	for item := range input {
		if !has || item < minItem {
			minItem = item
			has = true
		}
	}

	return minItem, has
}

// Min2 finds the minimum (key, value) pair in a sequence.
//
// It returns the minimum key, value, and a boolean indicating if a minimum was found.
func Min2[K cmp.Ordered, V any](input iter.Seq2[K, V]) (K, V, bool) {
	var (
		minKey K
		minVal V
	)

	has := false

	for key, val := range input {
		if !has || key < minKey {
			minKey = key
			minVal = val
			has = true
		}
	}

	return minKey, minVal, has
}
