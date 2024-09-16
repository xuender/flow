package seq

import "iter"

// Last returns the last element of the sequence if it exists.
//
// Iterates through the provided sequence and keeps track of the latest element.
// Returns the last element along with a boolean indicating if an element was found.
func Last[V any](input iter.Seq[V]) (V, bool) {
	var (
		ret V
		has bool
	)

	for item := range input {
		ret = item
		has = true
	}

	return ret, has
}

// Last2 returns the last element of the sequence if it exists.
//
// It iterates through the provided sequence and keeps track of the latest element.
// Returns the key and value of the last element along with a boolean indicating if an element was found.
func Last2[K, V any](input iter.Seq2[K, V]) (K, V, bool) {
	var (
		key K
		val V
		has bool
	)

	for itemKey, itemVal := range input {
		key = itemKey
		val = itemVal
		has = true
	}

	return key, val, has
}
