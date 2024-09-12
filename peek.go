package flow

import "iter"

func Peek[E any](action func(E)) Step[E, E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
		return func(yield func(E) bool) {
			for item := range input {
				action(item)

				if !yield(item) {
					return
				}
			}
		}
	}
}
