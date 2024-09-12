package flow

import (
	"cmp"
	"iter"
	"slices"
)

func Sort[E cmp.Ordered]() Step[E, E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
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
}
