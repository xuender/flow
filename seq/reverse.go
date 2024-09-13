package seq

import (
	"iter"
	"slices"
)

func Reverse[E any](input iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		slice := slices.Collect(input)
		slices.Reverse(slice)

		for _, item := range slice {
			if !yield(item) {
				return
			}
		}
	}
}
