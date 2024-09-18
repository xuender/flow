package seq

import "iter"

// TakeWhile filters elements from 'input' using 'predicate'.
//
// It yields elements until the predicate is false for the first time.
//
// It returned is a Seq[V] that can be used to iterate over the filtered elements.
//
// Play: https://go.dev/play/p/bNjM_sEmBVd
func TakeWhile[V any](input iter.Seq[V], predicate func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for item := range input {
			if !predicate(item) {
				return
			}

			if !yield(item) {
				return
			}
		}
	}
}

// TakeWhile2 filters elements from 'input' using 'predicate'.
//
// It yields key-value pairs until the predicate is false for the first time.
//
// It returned a Seq2[K, V] that can be used to iterate over the filtered key-value pairs.
//
// Play: https://go.dev/play/p/7ZvELBRdWSW
func TakeWhile2[K, V any](input iter.Seq2[K, V], predicate func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for key, val := range input {
			if !predicate(key, val) {
				return
			}

			if !yield(key, val) {
				return
			}
		}
	}
}

// DropWhile skips elements from 'input' as long as 'predicate' returns true.
//
// It returned function is a Seq[V] that can be used to iterate over the remaining elements.
//
// Play: https://go.dev/play/p/fAUtx4arGB7
func DropWhile[V any](input iter.Seq[V], predicate func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		pass := true

		for item := range input {
			if pass && predicate(item) {
				continue
			}

			pass = false

			if !yield(item) {
				return
			}
		}
	}
}

// DropWhile2 skips key-value pairs from 'input' as long as 'predicate' returns true.
//
// Once the predicate is false, it starts yielding pairs.
//
// It returned a Seq2[K, V] that can be used to iterate over the remaining pairs.
//
// Play: https://go.dev/play/p/V_LQKrSZFiF
func DropWhile2[K, V any](input iter.Seq2[K, V], predicate func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		pass := true

		for key, val := range input {
			if pass && predicate(key, val) {
				continue
			}

			pass = false

			if !yield(key, val) {
				return
			}
		}
	}
}
