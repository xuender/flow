package seq

import "iter"

// AnyMatch checks if any element in the input sequence satisfies the given predicate.
//
// This function iterates over the elements in `input` and applies the `predicate` function to each element.
// If any element satisfies the predicate, the function returns true; otherwise, it returns false.
//
// Args:
//
//	input (iter.Seq[E]): The input sequence of elements to check.
//	predicate (func(E) bool): A function that tests if an element satisfies the condition.
//
// Returns:
//
//	bool: True if any element satisfies the predicate, false otherwise.
func AnyMatch[E any](input iter.Seq[E], predicate func(E) bool) bool {
	for item := range input {
		if predicate(item) {
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
//	input (iter.Seq[E]): The input sequence of elements to check.
//	predicate (func(E) bool): A function that tests if an element satisfies the condition.
//
// Returns:
//
//	bool: True if all elements satisfy the predicate, false otherwise.
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

// NoneMatch checks if no elements in the input sequence satisfy the given predicate.
//
// This function uses `AnyMatch` to determine if any element satisfies the predicate.
// If no elements satisfy the predicate, the function returns true; otherwise, it returns false.
//
// Args:
//
//	input (iter.Seq[E]): The input sequence of elements to check.
//	predicate (func(E) bool): A function that tests if an element satisfies the condition.
//
// Returns:
//
//	bool: True if no elements satisfy the predicate, false otherwise.
func NoneMatch[E any](input iter.Seq[E], predicate func(E) bool) bool {
	return !AnyMatch(input, predicate)
}
