package seq

import "iter"

// AnyMatch checks if any element in the input sequence satisfies the given predicate.
//
// This function iterates over the elements in `input` and applies the `predicate` function to each element.
// If any element satisfies the predicate, the function returns true; otherwise, it returns false.
//
// Args:
//
//	input iter.Seq[V]: The input sequence of elements to check.
//	predicate func(V) bool: A function that tests if an element satisfies the condition.
//
// Returns:
//
//	bool: True if any element satisfies the predicate, false otherwise.
func AnyMatch[V any](input iter.Seq[V], predicate func(V) bool) bool {
	for item := range input {
		if predicate(item) {
			return true
		}
	}

	return false
}

// AnyMatch2 checks if any (key, value) pair in the sequence satisfies a predicate.
// It returns true if at least one pair satisfies the predicate.
//
// Args:
//
//	input iter.Seq2[K, V]: The input sequence of (key, value) pairs.
//	predicate func(K, V) bool: The predicate function to apply to each pair.
//
// Returns:
//
//	bool: True if any pair satisfies the predicate, false otherwise.
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
// This function iterates over the elements in `input` and applies the `predicate` function to each element.
// If all elements satisfy the predicate, the function returns true; otherwise, it returns false.
//
// Args:
//
//	input iter.Seq[V]: The input sequence of elements to check.
//	predicate func(V) bool: A function that tests if an element satisfies the condition.
//
// Returns:
//
//	bool: True if all elements satisfy the predicate, false otherwise.
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
// It returns true if all pairs satisfy the predicate.
//
// Args:
//
//	input iter.Seq2[K, V]: The input sequence of (key, value) pairs.
//	predicate func(K, V) bool: The predicate function to apply to each pair.
//
// Returns:
//
//	bool: True if all pairs satisfy the predicate, false otherwise.
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
// This function uses `AnyMatch` to determine if any element satisfies the predicate.
// If no elements satisfy the predicate, the function returns true; otherwise, it returns false.
//
// Args:
//
//	input iter.Seq[V]: The input sequence of elements to check.
//	predicate func(V) bool: A function that tests if an element satisfies the condition.
//
// Returns:
//
//	bool: True if no elements satisfy the predicate, false otherwise.
func NoneMatch[V any](input iter.Seq[V], predicate func(V) bool) bool {
	return !AnyMatch(input, predicate)
}

// NoneMatch2 checks if no (key, value) pairs in the sequence satisfy a predicate.
//
// It returns true if no pairs satisfy the predicate.
//
// Args:
//
//	input iter.Seq2[K, V]: The input sequence of (key, value) pairs.
//	predicate func(K, V) bool: The predicate function to apply to each pair.
//
// Returns:
//
//	bool: True if no pairs satisfy the predicate, false otherwise.
func NoneMatch2[K, V any](input iter.Seq2[K, V], predicate func(K, V) bool) bool {
	return !AnyMatch2(input, predicate)
}
