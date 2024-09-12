package flow

import "iter"

func Filter[E any](predicate func(E) bool) Step[E, E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
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
}
