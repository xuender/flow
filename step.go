package flow

import (
	"cmp"
	"iter"

	"gitee.com/xuender/flow/seq"
)

type Step[E any] func(iter.Seq[E]) iter.Seq[E]

func Filter[E any](predicate func(E) bool) Step[E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
		return seq.Filter(input, predicate)
	}
}

func Limit[E any](limit int) Step[E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
		return seq.Limit(input, limit)
	}
}

func Peek[E any](action func(E)) Step[E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
		return seq.Peek(input, action)
	}
}

func Reduce[E cmp.Ordered]() Step[E] {
	return seq.Reduce
}

func Reverse[E any]() Step[E] {
	return seq.Reverse
}

func Skip[E any](num int) Step[E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
		return seq.Skip(input, num)
	}
}

func Sort[E cmp.Ordered]() Step[E] {
	return seq.Sort
}

func SortFunc[E any](cmp func(item1, item2 E) int) Step[E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
		return seq.SortFunc(input, cmp)
	}
}

func Unique[E comparable]() Step[E] {
	return seq.Unique
}
