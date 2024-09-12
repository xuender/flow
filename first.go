package flow

import "iter"

func First[E any](seq iter.Seq[E]) (E, bool) {
	for item := range seq {
		return item, true
	}

	var zero E

	return zero, false
}
