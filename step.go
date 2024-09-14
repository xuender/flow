package flow

import (
	"cmp"
	"iter"

	"github.com/xuender/flow/seq"
)

// Step represents a transformation function that can be applied to a sequence.
//
// This type defines a function that takes an input sequence of elements and returns
// a new sequence after applying some transformation.
//
// Parameters:
//
//	E (any): The type of elements in the sequence.
//
// Returns:
//
//	func(iter.Seq[E]) iter.Seq[E]: A transformation function.
type Step[E any] func(iter.Seq[E]) iter.Seq[E]

// Distinct returns a transformation step that filters out duplicate elements from a sequence.
//
// This function returns a `Step` that can be used to create a new sequence containing only
// the distinct elements from the input sequence.
//
// Parameters:
//
//	E (comparable): The type of elements in the sequence.
//
// Returns:
//
//	Step[E]: A transformation step that filters duplicates.
func Distinct[E comparable]() Step[E] {
	return seq.Distinct
}

// Filter returns a transformation step that filters elements based on a predicate.
//
// This function returns a `Step` that can be used to create a new sequence containing only
// the elements that satisfy the given `predicate`.
//
// Parameters:
//
//	predicate (func(E) bool): The predicate function to filter elements.
//
// Returns:
//
//	Step[E]: A transformation step that filters elements.
func Filter[E any](predicate func(E) bool) Step[E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
		return seq.Filter(input, predicate)
	}
}

// Limit returns a transformation step that limits the number of elements in a sequence.
//
// This function returns a `Step` that creates a new sequence containing at most `limit` elements
// from the input sequence.
//
// Parameters:
//
//	limit (int): The maximum number of elements to include.
//
// Returns:
//
//	Step[E]: A transformation step that limits the sequence.
func Limit[E any](limit int) Step[E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
		return seq.Limit(input, limit)
	}
}

// Peek returns a transformation step that applies an action to each element in the sequence.
//
// This function returns a `Step` that calls the `action` function on each element of the input
// sequence without modifying the sequence itself.
//
// Parameters:
//
//	action (func(E)): The action to apply to each element.
//
// Returns:
//
//	Step[E]: A transformation step that applies the action.
func Peek[E any](action func(E)) Step[E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
		return seq.Peek(input, action)
	}
}

// Reverse returns a transformation step that reverses the order of elements in a sequence.
//
// This function returns a `Step` that creates a new sequence with the elements in reverse order.
//
// Returns:
//
//	Step[E]: A transformation step that reverses the sequence.
func Reverse[E any]() Step[E] {
	return seq.Reverse
}

// Skip returns a transformation step that skips the first `num` elements in a sequence.
//
// This function returns a `Step` that creates a new sequence by skipping the first `num` elements
// of the input sequence.
//
// Parameters:
//
//	num (int): The number of elements to skip.
//
// Returns:
//
//	Step[E]: A transformation step that skips elements.
func Skip[E any](num int) Step[E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
		return seq.Skip(input, num)
	}
}

// Sort returns a transformation step that sorts the elements in a sequence.
//
// This function returns a `Step` that creates a new sequence with the elements sorted.
//
// Returns:
//
//	Step[E]: A transformation step that sorts the sequence.
func Sort[E cmp.Ordered]() Step[E] {
	return seq.Sort
}

// SortFunc returns a transformation step that sorts the elements in a sequence using a custom comparison function.
//
// This function returns a `Step` that creates a new sequence with the elements sorted according to the `cmp` function.
//
// Parameters:
//
//	cmp (func(item1, item2 E) int): The comparison function for sorting.
//
// Returns:
//
//	Step[E]: A transformation step that sorts the sequence.
func SortFunc[E any](cmp func(item1, item2 E) int) Step[E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
		return seq.SortFunc(input, cmp)
	}
}
