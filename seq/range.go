package seq

import (
	"iter"
)

// Range generates a sequence of integers from start to end (exclusive) with a given step.
// It returns a function that can be used as an iterator.
//
// Args:
//
//	A variadic list of uint values representing the start, end, and step.
//	  0 arguments: Empty sequence.
//	  1 argument: Generates from 0 to args[0] (exclusive).
//	  2 arguments: Generates from args[0] to args[1] (exclusive).
//	  3 or more arguments: Generates from args[0] to args[1] (exclusive) with a step of args[2].
func Range(args ...int) iter.Seq[int] {
	return func(yield func(int) bool) {
		start, end, step := readArgs(args)

		for range rangeLen(start, end, step) {
			if !yield(start) {
				return
			}

			start += step
		}
	}
}

// Range2 generates a sequence of integers from start to end (exclusive) with a given step.
// It returns a function that can be used as an iterator, yielding both the index and value.
//
// Args:
//
//	A variadic list of int values representing the start, end, and step.
//	  0 arguments: Empty sequence.
//	  1 argument: Generates from 0 to args[0] (exclusive).
//	  2 arguments: Generates from args[0] to args[1] (exclusive).
//	  3 or more arguments: Generates from args[0] to args[1] (exclusive) with a step of args[2].
func Range2(args ...int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		start, end, step := readArgs(args)
		idx := 0

		for range rangeLen(start, end, step) {
			if !yield(idx, start) {
				return
			}

			start += step
			idx++
		}
	}
}

func rangeLen(start, end, step int) int {
	if start == end || (end > start && step < 0) || (end < start && step > 0) {
		return 0
	}

	num := step
	if num == 0 {
		num = 1
	}

	length := (end - start) / num
	if (end-start)%num != 0 {
		return length + 1
	}

	return length
}

func readArgs(args []int) (int, int, int) {
	two := 2

	switch len(args) {
	case 0:
		return 0, 0, 0
	case 1:
		if args[0] < 0 {
			return 0, args[0], -1
		}

		if args[0] > 0 {
			return 0, args[0], 1
		}

		return 0, 0, 0
	case two:
		if args[0] < args[1] {
			return args[0], args[1], 1
		}

		if args[0] > args[1] {
			return args[0], args[1], -1
		}

		return 0, 0, 0
	default:
		return args[0], args[1], args[2]
	}
}
