package seq

import (
	"cmp"
	"iter"
)

// Max function finds the maximum element in the given sequence 'input'.
// It returns the maximum element and a boolean indicating whether a maximum was found.
// This function works with elements of any ordered type E that implements the cmp.Ordered interface.
//
// Parameters:
//
//	input: The sequence to find the maximum element from, of type iter.Seq[E].
//
// Returns:
//
//	E: The maximum element found in the sequence.
//	bool: A boolean value indicating whether a maximum element was found (true) or the sequence was empty (false).
func Max[E cmp.Ordered](input iter.Seq[E]) (E, bool) {
	var maxItem E

	has := false

	for item := range input {
		if !has || item > maxItem {
			maxItem = item
			has = true
		}
	}

	return maxItem, has
}
