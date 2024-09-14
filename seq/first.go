package seq

import "iter"

// First returns the first element of the input sequence or false if the sequence is empty.
//
// This function iterates over the sequence `input` and returns the first element found.
// If the sequence is empty, it returns the zero value of type `E` and false.
//
// Parameters:
//
//	input (iter.Seq[E]): The input sequence of elements.
//
// Returns:
//
//	E: The first element of the sequence.
//	bool: True if an element is found, false if the sequence is empty.
func First[E any](input iter.Seq[E]) (E, bool) {
	for item := range input {
		return item, true
	}

	var zero E

	return zero, false
}
