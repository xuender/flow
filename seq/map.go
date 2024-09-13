package seq

import "iter"

func Map[I, O any](input iter.Seq[I], mapper func(I) O) iter.Seq[O] {
	return func(yield func(O) bool) {
		for item := range input {
			if !yield(mapper(item)) {
				return
			}
		}
	}
}
