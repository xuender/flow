package seq

import (
	"iter"
)

// Range generates a sequence of integers from start to end (exclusive) with a given step.
// It returns a function that can be used as an iterator.
//
// Parameters:
//
//	args: A variadic list of uint values representing the start, end, and step.
//	  0 arguments: Empty sequence.
//	  1 argument: Generates from 0 to args[0] (exclusive).
//	  2 arguments: Generates from args[0] to args[1] (exclusive).
//	  3 or more arguments: Generates from args[0] to args[1] (exclusive) with a step of args[2].
//
// Returns:
//
//	A function that yields integers in the specified range.
func Range(args ...int) iter.Seq[int] {
	return func(yield func(int) bool) {
		start, end, step := readArgs(args)
		if start == end || (end > start && step < 0) || (end < start && step > 0) {
			return
		}

		num := step
		if num == 0 {
			num = 1
		}

		length := (end - start) / num
		if (end-start)%num != 0 {
			length++
		}

		for range length {
			if !yield(start) {
				return
			}

			start += step
		}
	}
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
