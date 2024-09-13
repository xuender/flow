package flow

import (
	"cmp"
	"iter"

	"gitee.com/xuender/flow/seq"
)

type Step[I, O any] func(iter.Seq[I]) iter.Seq[O]

func Filter[E any](predicate func(E) bool) Step[E, E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
		return seq.Filter(input, predicate)
	}
}

func FlatMap[S ~[]E, E any]() Step[S, E] {
	return seq.FlatMap
}

func Limit[E any](limit int) Step[E, E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
		return seq.Limit(input, limit)
	}
}

func Map[I, O any](mappger func(I) O) Step[I, O] {
	return func(input iter.Seq[I]) iter.Seq[O] {
		return seq.Map(input, mappger)
	}
}

func Peek[E any](action func(E)) Step[E, E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
		return seq.Peek(input, action)
	}
}

func Reverse[E any]() Step[E, E] {
	return seq.Reverse
}

func Skip[E any](num int) Step[E, E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
		return seq.Skip(input, num)
	}
}

func Sort[E cmp.Ordered]() Step[E, E] {
	return seq.Sort
}

func Unique[E comparable]() Step[E, E] {
	return seq.Unique
}
