package flow

import "iter"

func Skip[E any](num int) Step[E, E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
		return func(yield func(E) bool) {
			idx := 0

			for item := range input {
				idx++
				if idx <= num {
					continue
				}

				if !yield(item) {
					return
				}
			}
		}
	}
}
