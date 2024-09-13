package seq

import "iter"

func Peek[E any](input iter.Seq[E], action func(E)) iter.Seq[E] {
	return func(yield func(E) bool) {
		for item := range input {
			action(item)

			if !yield(item) {
				return
			}
		}
	}
}
