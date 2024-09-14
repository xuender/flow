package seq

import (
	"cmp"
	"iter"
)

// Max returns the maximum element in the input sequence or false if the sequence is empty.
//
// This function iterates over the sequence `input` and finds the maximum element.
// If the sequence is empty, it returns the zero value of type `E` and false.
//
// Parameters:
//
//	input (iter.Seq[E]): The input sequence of ordered elements.
//
// Returns:
//
//	E: The maximum element in the sequence.
//	bool: True if an element is found, false if the sequence is empty.
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
