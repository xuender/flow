package seq

import "iter"

func Skip[E any](input iter.Seq[E], num int) iter.Seq[E] {
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
