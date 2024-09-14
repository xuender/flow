package seq

import "iter"

// Each iterates over the given sequence 'input' and executes the 'yield' function for each element.
// This function works with elements of any type E.
//
// Parameters:
//
//	input: The sequence to iterate over, of type iter.Seq[E].
//	yield: A function that takes an element as an argument and returns a boolean.
//	       If it returns false, the Each function will stop iterating and return.
//
// Returns:
//
//	No return value.
func Each[E any](input iter.Seq[E], yield func(E) bool) {
	for item := range input {
		if !yield(item) {
			return
		}
	}
}
