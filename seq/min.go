package seq

import (
	"cmp"
	"iter"
)

// Min returns the minimum element in the input sequence or false if the sequence is empty.
//
// This function iterates over the sequence `input` and finds the minimum element.
// If the sequence is empty, it returns the zero value of type `E` and false.
//
// Args:
//
//	input (iter.Seq[E]): The input sequence of ordered elements.
//
// Returns:
//
//	E: The minimum element in the sequence.
//	bool: True if an element is found, false if the sequence is empty.
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
