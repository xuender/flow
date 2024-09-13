package seq

import "iter"

func Filter[E any](input iter.Seq[E], predicate func(E) bool) iter.Seq[E] {
	return func(yield func(E) bool) {
		for item := range input {
			if predicate(item) {
				if !yield(item) {
					return
				}
			}
		}
	}
}
