package seq

import "iter"

func Emit[E any](input iter.Seq[E]) {
	input(func(E) bool {
		return true
	})
}
