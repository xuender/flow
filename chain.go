package flow

import "iter"

// Chain applies a series of transformation steps to the input sequence sequentially.
//
// This function takes an input sequence `input` and a variadic list of `steps`.
// Each step is applied sequentially to the input sequence.
//
// Args:
//
//	input iter.Seq[V]: The input sequence of elements.
//	steps ...Step[V]: A list of transformation steps to apply.
//
// Returns:
//
//	iter.Seq[V]: The transformed sequence after applying all steps.
func Chain[V any](input iter.Seq[V], steps ...Step[V]) iter.Seq[V] {
	for _, step := range steps {
		input = step(input)
	}

	return input
}

// Chain2 applies multiple Step2 functions sequentially to the input sequence.
//
// It returns the final sequence after applying all steps.
//
// Args:
//
//	input iter.Seq2[K, V]: The input sequence.
//	steps ...Step2[K, V]: The steps to apply.
//
// Returns:
//
//	iter.Seq2[K, V]: The final sequence.
func Chain2[K, V any](input iter.Seq2[K, V], steps ...Step2[K, V]) iter.Seq2[K, V] {
	for _, step := range steps {
		input = step(input)
	}

	return input
}
