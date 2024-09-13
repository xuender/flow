package seq

import (
	"cmp"
	"iter"
	"slices"
)

func Sort[E cmp.Ordered](input iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		slice := slices.Collect(input)
		slices.Sort(slice)

		for _, item := range slice {
			if !yield(item) {
				return
			}
		}
	}
}
