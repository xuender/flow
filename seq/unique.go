package seq

import (
	"iter"
)

func Unique[E comparable](input iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		none := struct{}{}
		set := map[E]struct{}{}

		for item := range input {
			if _, has := set[item]; has {
				continue
			}

			set[item] = none

			if !yield(item) {
				return
			}
		}
	}
}
