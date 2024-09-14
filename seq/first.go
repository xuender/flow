package seq

import "iter"

// First function retrieves the first element from the given sequence 'input'.
// If the sequence is empty, it returns the zero value of type E and false.
// This function works with elements of any type E.
//
// Parameters:
//
//	input: The sequence to retrieve the first element from, of type iter.Seq[E].
//
// Returns:
//
//	E: The first element of the sequence.
//	bool: A boolean value indicating whether an element was found (true) or the sequence was empty (false).
func First[E any](input iter.Seq[E]) (E, bool) {
	for item := range input {
		return item, true
	}

	var zero E

	return zero, false
}
