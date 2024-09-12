package flow

import (
	"iter"
	"slices"
)

func Reverse[E any]() Step[E, E] {
	return func(input iter.Seq[E]) iter.Seq[E] {
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
}
