package seq

import "iter"

func First[E any](input iter.Seq[E]) (E, bool) {
	for item := range input {
		return item, true
	}

	var zero E

	return zero, false
}
