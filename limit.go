package flow

import "iter"

func Limit[E any](size int) Step[E, E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
		return func(yield func(E) bool) {
			idx := 0

			for item := range input {
				if idx >= size {
					return
				}

				if !yield(item) {
					return
				}

				idx++
			}
		}
	}
}
