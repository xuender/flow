package seq

import "iter"

func Limit[E any](input iter.Seq[E], limit int) iter.Seq[E] {
	return func(yield func(E) bool) {
		idx := 0

		for item := range input {
			if idx >= limit {
				return
			}

			if !yield(item) {
				return
			}

			idx++
		}
	}
}
