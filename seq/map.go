package seq

import "iter"

// Map applies a transformation function to each element in the input sequence and returns a new sequence.
//
// This function takes a sequence `input` and a `mapper` function. It applies the `mapper` function
// to each element of the sequence and returns a new sequence with the transformed elements.
func Map[I, O any](input iter.Seq[I], mapper func(I) O) iter.Seq[O] {
	return func(yield func(O) bool) {
		for item := range input {
			if !yield(mapper(item)) {
				return
			}
		}
	}
}

// Map2 transforms a sequence of (key, value) pairs using a mapper function.
//
// It returns a new sequence with transformed pairs.
func Map2[IK, IV, OK, OV any](input iter.Seq2[IK, IV], mapper func(IK, IV) (OK, OV)) iter.Seq2[OK, OV] {
	return func(yield func(OK, OV) bool) {
		for key, val := range input {
			if !yield(mapper(key, val)) {
				return
			}
		}
	}
}
