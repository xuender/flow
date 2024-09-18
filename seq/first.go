package seq

import "iter"

// First returns the first element of the input sequence or false if the sequence is empty.
//
// It iterates over the sequence `input` and returns the first element found.
// If the sequence is empty, it returns the zero value of type `E` and false.
//
// Play: https://go.dev/play/p/sGB5Xd43vM3
func First[V any](input iter.Seq[V]) (V, bool) {
	for item := range input {
		return item, true
	}

	var zero V

	return zero, false
}

// First2 returns the first (key, value) pair in the sequence.
//
// It returns the first pair and a boolean indicating if a pair was found.
//
// Play: https://go.dev/play/p/hj0xJX0di1U
func First2[K, V any](input iter.Seq2[K, V]) (K, V, bool) {
	for key, val := range input {
		return key, val, true
	}

	var (
		key K
		val V
	)

	return key, val, false
}
