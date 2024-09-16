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
type Step[V any] func(iter.Seq[V]) iter.Seq[V]

// Step2 represents a function that processes a sequence of (key, value) pairs.
//
// It returns a new sequence of (key, value) pairs.
type Step2[K, V any] func(iter.Seq2[K, V]) iter.Seq2[K, V]

// Append adds multiple elements to the end of a sequence.
//
// It takes a variable number of elements and appends them to the input sequence.
func Append[V any](items ...V) Step[V] {
	return func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Append(input, items...)
	}
}

// Append2 appends additional (key, value) tuples to the input sequence.
//
// It returns a new sequence with the appended tuples.
func Append2[K, V any](items ...seq.Tuple[K, V]) Step2[K, V] {
	return func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Append2(input, items...)
	}
}

// Distinct returns a transformation step that filters out duplicate elements from a sequence.
//
// This function returns a `Step` that can be used to create a new sequence containing only
// the distinct elements from the input sequence.
func Distinct[V comparable]() Step[V] {
	return seq.Distinct
}

// Distinct2 removes duplicate (key, value) pairs from the sequence.
//
// It returns a Step2 function that filters duplicates.
func Distinct2[K comparable, V any]() Step2[K, V] {
	return seq.Distinct2
}

// Filter returns a transformation step that filters elements based on a predicate.
//
// This function returns a `Step` that can be used to create a new sequence containing only
// the elements that satisfy the given `predicate`.
func Filter[V any](predicate func(V) bool) Step[V] {
	return func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Filter(input, predicate)
	}
}

// Filter2 filters (key, value) pairs based on a predicate function.
//
// It returns a new sequence with filtered pairs.
func Filter2[K, V any](predicate func(K, V) bool) Step2[K, V] {
	return func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Filter2(input, predicate)
	}
}

// Limit returns a transformation step that limits the number of elements in a sequence.
//
// This function returns a `Step` that creates a new sequence containing at most `limit` elements
// from the input sequence.
func Limit[V any](limit int) Step[V] {
	return func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Limit(input, limit)
	}
}

// Limit2 limits the number of (key, value) pairs in the sequence.
//
// It returns a new sequence with a maximum number of pairs.
func Limit2[K, V any](limit int) Step2[K, V] {
	return func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Limit2(input, limit)
	}
}

// Merge combines multiple sequences into a single sequence.
//
// It merges the input sequence with additional sequences.
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
func Peek[V any](action func(V)) Step[V] {
	return func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Peek(input, action)
	}
}

// Peek2 applies an action to each (key, value) pair in the sequence.
//
// It returns a new sequence with the same pairs.
func Peek2[K, V any](action func(K, V)) Step2[K, V] {
	return func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Peek2(input, action)
	}
}

// Repeat creates a new step that repeats the input sequence a specified number of times.
func Repeat[V any](count int) Step[V] {
	return func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Repeat(input, count)
	}
}

// Repeat2 creates a new step that repeats each element of the input sequence a specified number of times.
//
// This function returns a `Step2` that repeats each element of the input sequence `count` times.
func Repeat2[K, V any](count int) Step2[K, V] {
	return func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Repeat2(input, count)
	}
}

// Reverse returns a transformation step that reverses the order of elements in a sequence.
//
// This function returns a `Step` that creates a new sequence with the elements in reverse order.
func Reverse[V any]() Step[V] {
	return seq.Reverse
}

// Reverse2 reverses the order of (key, value) pairs in the sequence.
//
// It returns a Step2 function that reverses the sequence.
func Reverse2[K, V any]() Step2[K, V] {
	return seq.Reverse2
}

// Skip returns a transformation step that skips the first `num` elements in a sequence.
//
// This function returns a `Step` that creates a new sequence by skipping the first `num` elements
// of the input sequence.
func Skip[V any](num int) Step[V] {
	return func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Skip(input, num)
	}
}

// Skip2 skips the first 'num' (key, value) pairs in the sequence.
//
// It returns a new sequence with the remaining pairs.
func Skip2[K, V any](num int) Step2[K, V] {
	return func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Skip2(input, num)
	}
}

// Sort returns a transformation step that sorts the elements in a sequence.
//
// This function returns a `Step` that creates a new sequence with the elements sorted.
func Sort[V cmp.Ordered]() Step[V] {
	return seq.Sorted
}

// Sort2 sorts the (key, value) pairs in the sequence.
//
// It returns a Step2 function that sorts the sequence.
func Sort2[K cmp.Ordered, V any]() Step2[K, V] {
	return seq.Sorted2
}

// SortFunc returns a transformation step that sorts the elements in a sequence using a custom comparison function.
//
// This function returns a `Step` that creates a new sequence with the elements sorted according to the `cmp` function.
func SortFunc[V any](cmp func(V, V) int) Step[V] {
	return func(input iter.Seq[V]) iter.Seq[V] {
		return seq.SortedFunc(input, cmp)
	}
}

// SortStableFunc sorts the elements in the sequence using a custom comparison function with stable sorting.
//
// It returns a Step function that sorts the sequence.
func SortStableFunc[V any](cmp func(V, V) int) Step[V] {
	return func(input iter.Seq[V]) iter.Seq[V] {
		return seq.SortedStableFunc(input, cmp)
	}
}

// SortFunc2 sorts (key, value) pairs using a custom comparison function.
//
// It returns a Step2 function that sorts the sequence.
func SortFunc2[K, V any](cmp func(seq.Tuple[K, V], seq.Tuple[K, V]) int) Step2[K, V] {
	return func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.SortedFunc2(input, cmp)
	}
}

// SortStableFunc2 sorts (key, value) pairs using a custom comparison function with stable sorting.
//
// It returns a Step2 function that sorts the sequence.
func SortStableFunc2[K, V any](cmp func(seq.Tuple[K, V], seq.Tuple[K, V]) int) Step2[K, V] {
	return func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.SortedStableFunc2(input, cmp)
	}
}
