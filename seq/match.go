package seq

import "iter"

// AnyMatch checks if any element in the input sequence satisfies the given predicate.
//
// If any element satisfies the predicate, the function returns true; otherwise, it returns false.
func AnyMatch[V any](input iter.Seq[V], predicate func(V) bool) bool {
	for item := range input {
		if predicate(item) {
			return true
		}
	}

	return false
}

// AnyMatch2 checks if any (key, value) pair in the sequence satisfies a predicate.
//
// It returns true if at least one pair satisfies the predicate.
func AnyMatch2[K, V any](input iter.Seq2[K, V], predicate func(K, V) bool) bool {
	for key, item := range input {
		if predicate(key, item) {
			return true
		}
	}

	return false
}

// AllMatch checks if all elements in the input sequence satisfy the given predicate.
//
// If all elements satisfy the predicate, the function returns true; otherwise, it returns false.
func AllMatch[V any](input iter.Seq[V], predicate func(V) bool) bool {
	ret := false

	for item := range input {
		if !predicate(item) {
			return false
		}

		ret = true
	}

	return ret
}

// AllMatch2 checks if all (key, value) pairs in the sequence satisfy a predicate.
//
// It returns true if all pairs satisfy the predicate.
func AllMatch2[K, V any](input iter.Seq2[K, V], predicate func(K, V) bool) bool {
	ret := false

	for key, val := range input {
		if !predicate(key, val) {
			return false
		}

		ret = true
	}

	return ret
}

// NoneMatch checks if no elements in the input sequence satisfy the given predicate.
//
// If no elements satisfy the predicate, the function returns true; otherwise, it returns false.
func NoneMatch[V any](input iter.Seq[V], predicate func(V) bool) bool {
	return !AnyMatch(input, predicate)
}

// NoneMatch2 checks if no (key, value) pairs in the sequence satisfy a predicate.
//
// It returns true if no pairs satisfy the predicate.
func NoneMatch2[K, V any](input iter.Seq2[K, V], predicate func(K, V) bool) bool {
	return !AnyMatch2(input, predicate)
}
