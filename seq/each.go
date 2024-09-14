package seq

import "iter"

func Each[E any](input iter.Seq[E], yield func(E) bool) {
	for item := range input {
		if !yield(item) {
			return
		}
	}
}
