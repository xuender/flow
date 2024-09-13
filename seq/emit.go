package seq

import "iter"

func Emit[E any](seq iter.Seq[E]) {
	call := func(E) bool {
		return true
	}

	seq(call)
}
