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
// Args:
//
//	V any: The type of elements in the sequence.
//
// Returns:
//
//	func(iter.Seq[V]) iter.Seq[V]: A transformation function.
type Step[V any] func(iter.Seq[V]) iter.Seq[V]

// Step2 represents a function that processes a sequence of (key, value) pairs.
//
// It returns a new sequence of (key, value) pairs.

// Args:
//
//	func(iter.Seq2[K, V]) iter.Seq2[K, V]: The function to process the sequence.
//
// Returns:
//
//	func(iter.Seq2[K, V]) iter.Seq2[K, V]: A transformation function.
type Step2[K, V any] func(iter.Seq2[K, V]) iter.Seq2[K, V]

// Append adds multiple elements to the end of a sequence.
//
// It takes a variable number of elements and appends them to the input sequence.
//
// Args:
//
//	items ...V: Elements to append to the sequence.
//
// Returns:
//
//	Step[V]: A function that takes an input sequence and returns a new sequence with the appended elements.
func Append[V any](items ...V) Step[V] {
	return func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Append(input, items...)
	}
}

// Append2 appends additional (key, value) tuples to the input sequence.
//
// It returns a new sequence with the appended tuples.
//
// Args:
//
//	items ...seq.Tuple[K, V]: The tuples to append.
//
// Returns:
//
//	Step2[K, V]: A function that appends tuples to the sequence.
func Append2[K, V any](items ...seq.Tuple[K, V]) Step2[K, V] {
	return func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Append2(input, items...)
	}
}

// Distinct returns a transformation step that filters out duplicate elements from a sequence.
//
// This function returns a `Step` that can be used to create a new sequence containing only
// the distinct elements from the input sequence.
//
// Args:
//
//	V (comparable): The type of elements in the sequence.
//
// Returns:
//
//	Step[V]: A transformation step that filters duplicates.
func Distinct[V comparable]() Step[V] {
	return seq.Distinct
}

// Distinct2 removes duplicate (key, value) pairs from the sequence.
//
// It returns a Step2 function that filters duplicates.
//
// Returns:
//
//	Step2[K, V]: A function that removes duplicate pairs.
func Distinct2[K comparable, V any]() Step2[K, V] {
	return seq.Distinct2
}

// Filter returns a transformation step that filters elements based on a predicate.
//
// This function returns a `Step` that can be used to create a new sequence containing only
// the elements that satisfy the given `predicate`.
// Args:
//
//	predicate func(V) bool: The predicate function to filter elements.
//
// Returns:
//
//	Step[V]: A transformation step that filters elements.
func Filter[V any](predicate func(V) bool) Step[V] {
	return func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Filter(input, predicate)
	}
}

// Filter2 filters (key, value) pairs based on a predicate function.
//
// It returns a new sequence with filtered pairs.
//
// Args:
//
//	predicate func(K, V) bool: The predicate function to filter pairs.
//
// Returns:
//
//	Step2[K, V]: A function that filters the sequence.
func Filter2[K, V any](predicate func(K, V) bool) Step2[K, V] {
	return func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Filter2(input, predicate)
	}
}

// Limit returns a transformation step that limits the number of elements in a sequence.
//
// This function returns a `Step` that creates a new sequence containing at most `limit` elements
// from the input sequence.
//
// Args:
//
//	limit int: The maximum number of elements to include.
//
// Returns:
//
//	Step[V]: A transformation step that limits the sequence.
func Limit[V any](limit int) Step[V] {
	return func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Limit(input, limit)
	}
}

// Limit2 limits the number of (key, value) pairs in the sequence.
//
// It returns a new sequence with a maximum number of pairs.
//
// Args:
//
//	limit int: The maximum number of pairs.
//
// Returns:
//
//	Step2[K, V]: A function that limits the sequence.
func Limit2[K, V any](limit int) Step2[K, V] {
	return func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Limit2(input, limit)
	}
}

// Merge combines multiple sequences into a single sequence.
//
// It merges the input sequence with additional sequences.
//
// Args:
//
//	seqs ...iter.Seq[V]: Additional sequences to merge.
//
// Returns:
//
//	Step[V]: A function that merges the input sequence with the additional sequences.
func Merge[V any](seqs ...iter.Seq[V]) Step[V] {
	return func(input iter.Seq[V]) iter.Seq[V] {
		if len(seqs) == 0 {
			return input
		}

		return seq.Merge(append([]iter.Seq[V]{input}, seqs...)...)
	}
}

// Merge2 merges multiple sequences of (key, value) pairs.
//
// It returns a new sequence containing all pairs.
//
// Args:
//
//	seqs ...iter.Seq2[K, V]: The sequences to merge.
//
// Returns:
//
//	Step2[K, V]: A function that merges the sequences.
func Merge2[K, V any](seqs ...iter.Seq2[K, V]) Step2[K, V] {
	return func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		if len(seqs) == 0 {
			return input
		}

		return seq.Merge2(append([]iter.Seq2[K, V]{input}, seqs...)...)
	}
}

// Peek returns a transformation step that applies an action to each element in the sequence.
//
// This function returns a `Step` that calls the `action` function on each element of the input
// sequence without modifying the sequence itself.
//
// Args:
//
//	action func(V): The action to apply to each element.
//
// Returns:
//
//	Step[V]: A transformation step that applies the action.
func Peek[V any](action func(V)) Step[V] {
	return func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Peek(input, action)
	}
}

// Peek2 applies an action to each (key, value) pair in the sequence.
//
// It returns a new sequence with the same pairs.
//
// Args:
//
//	action func(K, V): The action to apply to each pair.
//
// Returns:
//
//	Step2[K, V]: A function that applies the action.
func Peek2[K, V any](action func(K, V)) Step2[K, V] {
	return func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Peek2(input, action)
	}
}

// Reverse returns a transformation step that reverses the order of elements in a sequence.
//
// This function returns a `Step` that creates a new sequence with the elements in reverse order.
//
// Returns:
//
//	Step[V]: A transformation step that reverses the sequence.
func Reverse[V any]() Step[V] {
	return seq.Reverse
}

// Reverse2 reverses the order of (key, value) pairs in the sequence.
//
// It returns a Step2 function that reverses the sequence.
//
// Returns:
//
//	Step2[K, V]: A function that reverses the sequence.
func Reverse2[K, V any]() Step2[K, V] {
	return seq.Reverse2
}

// Skip returns a transformation step that skips the first `num` elements in a sequence.
//
// This function returns a `Step` that creates a new sequence by skipping the first `num` elements
// of the input sequence.
//
// Args:
//
//	num int: The number of elements to skip.
//
// Returns:
//
//	Step[V]: A transformation step that skips elements.
func Skip[V any](num int) Step[V] {
	return func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Skip(input, num)
	}
}

// Skip2 skips the first 'num' (key, value) pairs in the sequence.
//
// It returns a new sequence with the remaining pairs.
//
// Args:
//
//	num int: The number of pairs to skip.
//
// Returns:
//
//	Step2[K, V]: A function that skips pairs.
func Skip2[K, V any](num int) Step2[K, V] {
	return func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Skip2(input, num)
	}
}

// Sort returns a transformation step that sorts the elements in a sequence.
//
// This function returns a `Step` that creates a new sequence with the elements sorted.
//
// Returns:
//
//	Step[V]: A transformation step that sorts the sequence.
func Sort[V cmp.Ordered]() Step[V] {
	return seq.Sorted
}

// Sort2 sorts the (key, value) pairs in the sequence.
//
// It returns a Step2 function that sorts the sequence.
//
// Returns:
//
//	Step2[K, V]: A function that sorts the sequence.
func Sort2[K cmp.Ordered, V any]() Step2[K, V] {
	return seq.Sorted2
}

// SortFunc returns a transformation step that sorts the elements in a sequence using a custom comparison function.
//
// This function returns a `Step` that creates a new sequence with the elements sorted according to the `cmp` function.
//
// Args:
//
//	cmp func(V, V) int: The comparison function for sorting.
//
// Returns:
//
//	Step[V]: A transformation step that sorts the sequence.
func SortFunc[V any](cmp func(V, V) int) Step[V] {
	return func(input iter.Seq[V]) iter.Seq[V] {
		return seq.SortedFunc(input, cmp)
	}
}

// SortStableFunc sorts the elements in the sequence using a custom comparison function with stable sorting.
//
// It returns a Step function that sorts the sequence.
//
// Args:
//
//	cmp func(V, V) int: The comparison function.
//
// Returns:
//
//	Step[V]: A function that sorts the sequence.
func SortStableFunc[V any](cmp func(V, V) int) Step[V] {
	return func(input iter.Seq[V]) iter.Seq[V] {
		return seq.SortedStableFunc(input, cmp)
	}
}

// SortFunc2 sorts (key, value) pairs using a custom comparison function.
//
// It returns a Step2 function that sorts the sequence.
//
// Args:
//
//	cmp func(seq.Tuple[K, V], seq.Tuple[K, V]) int: The comparison function.
//
// Returns:
//
//	Step2[K, V]: A function that sorts the sequence.
func SortFunc2[K, V any](cmp func(seq.Tuple[K, V], seq.Tuple[K, V]) int) Step2[K, V] {
	return func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.SortedFunc2(input, cmp)
	}
}

// SortStableFunc2 sorts (key, value) pairs using a custom comparison function with stable sorting.
//
// It returns a Step2 function that sorts the sequence.
//
// Args:
//
//	cmp func(seq.Tuple[K, V], seq.Tuple[K, V]) int: The comparison function.
//
// Returns:
//
//	Step2[K, V]: A function that sorts the sequence.
func SortStableFunc2[K, V any](cmp func(seq.Tuple[K, V], seq.Tuple[K, V]) int) Step2[K, V] {
	return func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.SortedStableFunc2(input, cmp)
	}
}
