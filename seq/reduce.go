package seq

import (
	"cmp"
	"iter"
	"slices"
)

func Reduce[E cmp.Ordered](input iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		slice := slices.Collect(input)
		slices.Sort(slice)
		slices.Reverse(slice)

		for _, item := range slice {
			if !yield(item) {
				return
			}
		}
	}
}
