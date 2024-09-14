package seq

import (
	"iter"
)

// Reduce reduces the elements of the input sequence using the provided accumulator function.
//
// This function takes a sequence `input` and an `accumulator` function. It applies the accumulator
// function cumulatively to the items of the sequence, from left to right, to reduce the sequence.
//
// Args:
//
//	input iter.Seq[V]: The input sequence of elements.
//	accumulator func(V, V) V: The accumulator function to reduce the sequence.
//
// Returns:
//
//	V: The reduced result of the sequence.
//	bool: True if the reduction was performed, false if the sequence is empty.
func Reduce[V any](input iter.Seq[V], accumulator func(V, V) V) (V, bool) {
	var (
		ret      V
		isReduce bool
	)

	for item := range input {
		if !isReduce {
			ret = item
			isReduce = true

			continue
		}

		ret = accumulator(ret, item)
	}

	return ret, isReduce
}

// Reduce2 reduces a sequence of (key, value) pairs using an accumulator function.
//
// It returns the reduced result and a boolean indicating if reduction occurred.
//
// Args:
//
//	input iter.Seq2[K, V]: The input sequence of (key, value) pairs.
//	accumulator func(K, V, K, V) (K, V): The accumulator function to reduce pairs.
//
// Returns:
//
//	K: The reduced key.
//	V: The reduced value.
//	bool: Indicates if reduction occurred.
func Reduce2[K, V any](input iter.Seq2[K, V], accumulator func(K, V, K, V) (K, V)) (K, V, bool) {
	var (
		retKey   K
		retVal   V
		isReduce bool
	)

	for key, val := range input {
		if !isReduce {
			retKey = key
			retVal = val
			isReduce = true

			continue
		}

		retKey, retVal = accumulator(retKey, retVal, key, val)
	}

	return retKey, retVal, isReduce
}
