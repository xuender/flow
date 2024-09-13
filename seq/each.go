package seq

import "iter"

func Each[E any](seq iter.Seq[E], yield func(E) bool) {
	for item := range seq {
		if !yield(item) {
			return
		}
	}
}
