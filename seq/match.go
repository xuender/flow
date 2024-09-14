package seq

import "iter"

func AnyMatch[E any](input iter.Seq[E], predicate func(E) bool) bool {
	for item := range input {
		if predicate(item) {
			return true
		}
	}

	return false
}

func AllMatch[E any](input iter.Seq[E], predicate func(E) bool) bool {
	ret := false

	for item := range input {
		if !predicate(item) {
			return false
		}

		ret = true
	}

	return ret
}

func NoneMatch[E any](input iter.Seq[E], predicate func(E) bool) bool {
	return !AnyMatch(input, predicate)
}
