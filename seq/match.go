package seq

import "iter"

func AnyMatch[E any](seq iter.Seq[E], predicate func(E) bool) bool {
	for item := range seq {
		if predicate(item) {
			return true
		}
	}

	return false
}

func AllMatch[E any](seq iter.Seq[E], predicate func(E) bool) bool {
	ret := false

	for item := range seq {
		if !predicate(item) {
			return false
		}

		ret = true
	}

	return ret
}

func NoneMatch[E any](seq iter.Seq[E], predicate func(E) bool) bool {
	return !AnyMatch(seq, predicate)
}
