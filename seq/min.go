package seq

import (
	"cmp"
	"iter"
)

// Min function finds the minimum element in the given sequence 'input'.
// It returns the minimum element and a boolean indicating whether a minimum was found.
// This function works with elements of any ordered type E that implements the cmp.Ordered interface.
//
// Parameters:
//
//	input: The sequence to find the minimum element from, of type iter.Seq[E].
//
// Returns:
//
//	E: The minimum element found in the sequence.
//	bool: A boolean value indicating whether a minimum element was found (true) or the sequence was empty (false).
func Min[E cmp.Ordered](input iter.Seq[E]) (E, bool) {
	var minItem E

	has := false

	for item := range input {
		if !has || item < minItem {
			minItem = item
			has = true
		}
	}

	return minItem, has
}
