package flow

import "iter"

func Map[I, O any](mapper func(I) O) Step[I, O] {
	return func(input iter.Seq[I]) iter.Seq[O] {
		return func(yield func(O) bool) {
			for item := range input {
				if !yield(mapper(item)) {
					return
				}
			}
		}
	}
}
