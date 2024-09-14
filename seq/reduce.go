package seq

import (
	"iter"
)

// Reduce reduces the elements of the input sequence using the provided accumulator function.
//
// This function takes a sequence `input` and an `accumulator` function. It applies the accumulator
// function cumulatively to the items of the sequence, from left to right, to reduce the sequence.
//
// Args:
//
//	input (iter.Seq[E]): The input sequence of elements.
//	accumulator (func(E, E) E): The accumulator function to reduce the sequence.
//
// Returns:
//
//	E: The reduced result of the sequence.
//	bool: True if the reduction was performed, false if the sequence is empty.
func Reduce[E any](input iter.Seq[E], accumulator func(E, E) E) (E, bool) {
	var (
		ret      E
		isReduce bool
	)

	for item := range input {
		if !isReduce {
			ret = item
			isReduce = true

			continue
		}

		ret = accumulator(ret, item)
	}

	return ret, isReduce
}
