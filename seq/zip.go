package seq

import "iter"

// Zip combines elements from two sequences into pairs.
//
// It yields pairs of elements from `keys` and `values` until the shorter sequence is exhausted.
func Zip[K, V any](keys iter.Seq[K], values iter.Seq[V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		next, stop := iter.Pull(values)

		val, has := next()
		defer func() {
			if has {
				stop()
			}
		}()

		for key := range keys {
			if !yield(key, val) {
				return
			}

			if has {
				val, has = next()
			}
		}
	}
}
