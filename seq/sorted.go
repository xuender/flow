package seq

import (
	"cmp"
	"iter"
	"maps"
	"slices"
)

// Sorted sorts a sequence of elements that implement the cmp.Ordered interface.
//
// It returns a new sorted sequence.
//
// Play: https://go.dev/play/p/5Oe9nuwtJ8x
func Sorted[V cmp.Ordered](input iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, item := range slices.Sorted(input) {
			if !yield(item) {
				return
			}
		}
	}
}

// Sorted2 sorts the input sequence by keys and returns a sorted sequence.
//
// It takes a sequence of (key, value) pairs where keys are ordered.
//
// Play: https://go.dev/play/p/_msKNYeyauI
func Sorted2[K cmp.Ordered, V any](input iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		items := maps.Collect(input)

		for key := range Sorted(maps.Keys(items)) {
			if !yield(key, items[key]) {
				return
			}
		}
	}
}

// SortedFunc sorts the elements of the input sequence using a custom comparison function.
//
// It returns a new sorted sequence.
func SortedFunc[V any](input iter.Seq[V], cmp func(V, V) int) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, item := range slices.SortedFunc(input, cmp) {
			if !yield(item) {
				return
			}
		}
	}
}

// SortedStableFunc sorts the elements in the sequence using a custom comparison function with stable sorting.
//
// It returns a new sorted sequence.
func SortedStableFunc[V any](input iter.Seq[V], cmp func(V, V) int) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, item := range slices.SortedStableFunc(input, cmp) {
			if !yield(item) {
				return
			}
		}
	}
}

// SortedFunc2 sorts (key, value) pairs in the sequence using a custom comparison function.
//
// It returns a new sorted sequence.
func SortedFunc2[K, V any](input iter.Seq2[K, V], cmp func(Tuple[K, V], Tuple[K, V]) int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		slice := Tuples(input)

		slices.SortFunc(slice, cmp)

		for _, item := range slice {
			if !yield(item.K, item.V) {
				return
			}
		}
	}
}

// SortedStableFunc2 sorts (key, value) pairs using a custom comparison function with stable sorting.
//
// It returns a new sorted sequence.
func SortedStableFunc2[K, V any](input iter.Seq2[K, V], cmp func(Tuple[K, V], Tuple[K, V]) int) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		slice := Tuples(input)

		slices.SortStableFunc(slice, cmp)

		for _, item := range slice {
			if !yield(item.K, item.V) {
				return
			}
		}
	}
}
