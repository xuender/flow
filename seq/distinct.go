package seq

import (
	"iter"
)

// Distinct returns a new sequence containing only the distinct elements from the input sequence.
//
// This function iterates over the sequence `input` and filters out duplicate elements.
// Only the first occurrence of each element is included in the new sequence.
//
// Args:
//
//	input (iter.Seq[V]): The input sequence of comparable elements.
//
// Returns:
//
//	iter.Seq[V]: A new sequence with distinct elements.
func Distinct[V comparable](input iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		none := struct{}{}
		set := map[V]struct{}{}

		for item := range input {
			if _, has := set[item]; has {
				continue
			}

			set[item] = none

			if !yield(item) {
				return
			}
		}
	}
}

// Distinct2 returns a sequence that yields unique (key, value) pairs from the input sequence.
//
// It filters out duplicate keys based on their comparability.
//
// Args:
//
//	input iter.Seq2[K, V]: The input sequence of (key, value) pairs.
//
// Returns:
//
//	iter.Seq2[K, V]: A new sequence yielding only unique (key, value) pairs.
func Distinct2[K comparable, V any](input iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		none := struct{}{}
		set := map[K]struct{}{}

		for key, val := range input {
			if _, has := set[key]; has {
				continue
			}

			set[key] = none

			if !yield(key, val) {
				return
			}
		}
	}
}
