package seq

import (
	"iter"
)

// Reduce reduces the elements of the input sequence using the provided accumulator function.
//
// It returns the reduced result and a boolean indicating if reduction occurred.
//
// Play: https://go.dev/play/p/hnzytOtLAMi
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
// Play: https://go.dev/play/p/En1xY_ONKGO
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
