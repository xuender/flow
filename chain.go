package flow

import "iter"

// Chain applies a series of transformation steps to the input sequence sequentially.
//
// This function takes an input sequence `input` and a variadic list of `steps`.
// Each step is applied sequentially to the input sequence.
//
// Parameters:
//
//	input (iter.Seq[E]): The input sequence of elements.
//	steps (...Step[E]): A list of transformation steps to apply.
//
// Returns:
//
//	iter.Seq[E]: The transformed sequence after applying all steps.
func Chain[E any](input iter.Seq[E], steps ...Step[E]) iter.Seq[E] {
	for _, step := range steps {
		input = step(input)
	}

	return input
}
