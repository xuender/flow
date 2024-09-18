package flow

import (
	"cmp"
	"iter"

	"github.com/xuender/flow/seq"
)

// Step represents a processing step in an iterator.
//
// It includes a Next function to perform the next operation on a sequence,
// and a Parallel boolean indicating whether it can be executed in parallel.
type Step[V any] struct {
	Next     func(iter.Seq[V]) iter.Seq[V]
	Parallel bool
}

// Step2 represents a processing step in an iterator for sequences with two types K and V.
//
// It includes a Next function to perform the next operation on a sequence,
// and a Parallel boolean indicating whether it can be executed in parallel.
type Step2[K, V any] struct {
	Next     func(iter.Seq2[K, V]) iter.Seq2[K, V]
	Parallel bool
}

// Append adds multiple elements to the end of a sequence.
//
// It takes a variable number of elements and appends them to the input sequence.
func Append[V any](items ...V) Step[V] {
	return Step[V]{func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Append(input, items...)
	}, false}
}

// Append2 appends additional (key, value) tuples to the input sequence.
//
// It returns a new sequence with the appended tuples.
func Append2[K, V any](items ...seq.Tuple[K, V]) Step2[K, V] {
	return Step2[K, V]{func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Append2(input, items...)
	}, false}
}

// Distinct returns a transformation step that filters out duplicate elements from a sequence.
//
// It returns a `Step` that can be used to create a new sequence containing only
// the distinct elements from the input sequence.
func Distinct[V comparable]() Step[V] {
	return Step[V]{seq.Distinct[V], false}
}

// Distinct2 removes duplicate (key, value) pairs from the sequence.
//
// It returns a `Step2` function that filters duplicates.
func Distinct2[K comparable, V any]() Step2[K, V] {
	return Step2[K, V]{seq.Distinct2[K, V], false}
}

// DropWhile creates a Step[V] that skips elements from 'input' as long as 'predicate' returns true.
//
// Once the predicate is false, it starts yielding elements.

// It returned Step[V] can be used in a pipeline to process sequences.
func DropWhile[V any](predicate func(V) bool) Step[V] {
	return Step[V]{func(input iter.Seq[V]) iter.Seq[V] {
		return seq.DropWhile(input, predicate)
	}, false}
}

// DropWhile2 creates a Step2[K, V] that skips key-value pairs from 'input' as long as 'predicate' returns true.
//
// Once the predicate is false, it starts yielding pairs.
//
// It returned Step2[K, V] can be used in a pipeline to process sequences of key-value pairs.
func DropWhile2[K, V any](predicate func(K, V) bool) Step2[K, V] {
	return Step2[K, V]{func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.DropWhile2(input, predicate)
	}, false}
}

// Filter returns a transformation step that filters elements based on a predicate.
//
// It returns a `Step` that can be used to create a new sequence containing only
// the elements that satisfy the given `predicate`.
//
// Play: https://go.dev/play/p/JydmjWYw9rw
func Filter[V any](predicate func(V) bool) Step[V] {
	return Step[V]{func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Filter(input, predicate)
	}, true}
}

// Filter2 filters (key, value) pairs based on a predicate function.
//
// It returns a new sequence with filtered pairs.
func Filter2[K, V any](predicate func(K, V) bool) Step2[K, V] {
	return Step2[K, V]{func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Filter2(input, predicate)
	}, true}
}

// Limit returns a transformation step that limits the number of elements in a sequence.
//
// It returns a `Step` that creates a new sequence containing at most `limit` elements
// from the input sequence.
//
// Play: https://go.dev/play/p/JydmjWYw9rw
func Limit[V any](limit int) Step[V] {
	return Step[V]{func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Limit(input, limit)
	}, false}
}

// Limit2 limits the number of (key, value) pairs in the sequence.
//
// It returns a new sequence with a maximum number of pairs.
func Limit2[K, V any](limit int) Step2[K, V] {
	return Step2[K, V]{func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Limit2(input, limit)
	}, false}
}

// Map creates a new step that applies a transformation function to each element of an input sequence.
//
// It returns a Step that can be used in a pipeline to map over the input sequence.
//
// Note:
//
//	Cannot change type.
func Map[V any](maper func(V) V) Step[V] {
	return Step[V]{func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Map(input, maper)
	}, true}
}

// Map2 creates a new step that applies a transformation function to each pair of elements in an input sequence.
//
// It returns a Step2 that can be used in a pipeline to map over pairs in the input sequence.
//
// Note:
//
//	Cannot change type.
func Map2[K, V any](maper func(K, V) (K, V)) Step2[K, V] {
	return Step2[K, V]{func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Map2(input, maper)
	}, true}
}

// Merge combines multiple sequences into a single sequence.
//
// It merges the input sequence with additional sequences.
func Merge[V any](seqs ...iter.Seq[V]) Step[V] {
	return Step[V]{func(input iter.Seq[V]) iter.Seq[V] {
		if len(seqs) == 0 {
			return input
		}

		return seq.Merge(append([]iter.Seq[V]{input}, seqs...)...)
	}, false}
}

// Merge2 merges multiple sequences of (key, value) pairs.
//
// It returns a new sequence containing all pairs.
func Merge2[K, V any](seqs ...iter.Seq2[K, V]) Step2[K, V] {
	return Step2[K, V]{func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		if len(seqs) == 0 {
			return input
		}

		return seq.Merge2(append([]iter.Seq2[K, V]{input}, seqs...)...)
	}, false}
}

// Peek returns a transformation step that applies an action to each element in the sequence.
//
// It returns a `Step` that calls the `action` function on each element of the input
// sequence without modifying the sequence itself.
//
// Play: https://go.dev/play/p/0YiWlMdzwMk
func Peek[V any](action func(V)) Step[V] {
	return Step[V]{func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Peek(input, action)
	}, true}
}

// Prepend creates a step that adds items to the beginning of a sequence.
//
// It function is useful for creating a prepending step in a pipeline.
func Prepend[V any](items ...V) Step[V] {
	return Step[V]{func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Prepend(input, items...)
	}, true}
}

// Peek2 applies an action to each (key, value) pair in the sequence.
//
// It returns a new sequence with the same pairs.
func Peek2[K, V any](action func(K, V)) Step2[K, V] {
	return Step2[K, V]{func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Peek2(input, action)
	}, true}
}

// Prepend2 creates a step that adds tuples to the beginning of a sequence.
//
// It function is useful for creating a prepending step in a pipeline for tuples.
func Prepend2[K, V any](items ...seq.Tuple[K, V]) Step2[K, V] {
	return Step2[K, V]{func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Prepend2(input, items...)
	}, true}
}

// Repeat creates a new step that repeats the input sequence a specified number of times.
func Repeat[V any](count int) Step[V] {
	return Step[V]{func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Repeat(input, count)
	}, false}
}

// Repeat2 creates a new step that repeats each element of the input sequence a specified number of times.
//
// Tt returns a `Step2` that repeats each element of the input sequence `count` times.
func Repeat2[K, V any](count int) Step2[K, V] {
	return Step2[K, V]{func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Repeat2(input, count)
	}, false}
}

// Reverse returns a transformation step that reverses the order of elements in a sequence.
//
// It returns a `Step` that creates a new sequence with the elements in reverse order.
//
// Play: https://go.dev/play/p/JydmjWYw9rw
func Reverse[V any]() Step[V] {
	return Step[V]{seq.Reverse[V], false}
}

// Reverse2 reverses the order of (key, value) pairs in the sequence.
//
// It returns a `Step2` function that reverses the sequence.
func Reverse2[K, V any]() Step2[K, V] {
	return Step2[K, V]{seq.Reverse2[K, V], false}
}

// Shuffle creates a step that shuffles the input sequence.
//
// It returns a Step that, when applied, yields elements in a random order.
func Shuffle[V any]() Step[V] {
	return Step[V]{seq.Shuffle[V], false}
}

// Shuffle2 creates a step that shuffles pairs in the input sequence.
//
// It returns a Step2 that, when applied, yields pairs in a random order.
func Shuffle2[K, V any]() Step2[K, V] {
	return Step2[K, V]{seq.Shuffle2[K, V], false}
}

// Skip returns a transformation step that skips the first `num` elements in a sequence.
//
// It returns a `Step` that creates a new sequence by skipping the first `num` elements
// of the input sequence.
//
// Play: https://go.dev/play/p/JydmjWYw9rw
func Skip[V any](num int) Step[V] {
	return Step[V]{func(input iter.Seq[V]) iter.Seq[V] {
		return seq.Skip(input, num)
	}, false}
}

// Skip2 skips the first 'num' (key, value) pairs in the sequence.
//
// It returns a `Skip2` with the remaining pairs.
func Skip2[K, V any](num int) Step2[K, V] {
	return Step2[K, V]{func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.Skip2(input, num)
	}, false}
}

// Sort returns a transformation step that sorts the elements in a sequence.
//
// It returns a `Step` function that sorts the sequence.
func Sort[V cmp.Ordered]() Step[V] {
	return Step[V]{seq.Sorted[V], false}
}

// Sort2 sorts the (key, value) pairs in the sequence.
//
// It returns a `Step2` function that sorts the sequence.
func Sort2[K cmp.Ordered, V any]() Step2[K, V] {
	return Step2[K, V]{seq.Sorted2[K, V], false}
}

// SortFunc returns a transformation step that sorts the elements in a sequence using a custom comparison function.
//
// It returns a `Step` that creates a new sequence with the elements sorted according to the `cmp` function.
func SortFunc[V any](cmp func(V, V) int) Step[V] {
	return Step[V]{func(input iter.Seq[V]) iter.Seq[V] {
		return seq.SortedFunc(input, cmp)
	}, false}
}

// SortStableFunc sorts the elements in the sequence using a custom comparison function with stable sorting.
//
// It returns a `Step` function that sorts the sequence.
func SortStableFunc[V any](cmp func(V, V) int) Step[V] {
	return Step[V]{func(input iter.Seq[V]) iter.Seq[V] {
		return seq.SortedStableFunc(input, cmp)
	}, false}
}

// SortFunc2 sorts (key, value) pairs using a custom comparison function.
//
// It returns a `Step2` function that sorts the sequence.
func SortFunc2[K, V any](cmp func(seq.Tuple[K, V], seq.Tuple[K, V]) int) Step2[K, V] {
	return Step2[K, V]{func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.SortedFunc2(input, cmp)
	}, false}
}

// SortStableFunc2 sorts (key, value) pairs using a custom comparison function with stable sorting.
//
// It returns a `Step2` function that sorts the sequence.
func SortStableFunc2[K, V any](cmp func(seq.Tuple[K, V], seq.Tuple[K, V]) int) Step2[K, V] {
	return Step2[K, V]{func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.SortedStableFunc2(input, cmp)
	}, false}
}

// TakeWhile creates a Step[V] that filters elements from 'input' using 'predicate'.
//
// It yields elements until the predicate is false for the first time.
//
// The returned Step[V] can be used in a pipeline to process sequences.
func TakeWhile[V any](predicate func(V) bool) Step[V] {
	return Step[V]{func(input iter.Seq[V]) iter.Seq[V] {
		return seq.TakeWhile(input, predicate)
	}, false}
}

// TakeWhile2 creates a Step2[K, V] that filters key-value pairs from 'input' using 'predicate'.
//
// It yields pairs until the predicate is false for the first time.
//
// The returned Step2[K, V] can be used in a pipeline to process sequences of key-value pairs.
func TakeWhile2[K, V any](predicate func(K, V) bool) Step2[K, V] {
	return Step2[K, V]{func(input iter.Seq2[K, V]) iter.Seq2[K, V] {
		return seq.TakeWhile2(input, predicate)
	}, false}
}
