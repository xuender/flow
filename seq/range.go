package seq

import (
	"iter"
)

func Range(size int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for num := range size {
			if !yield(num) {
				return
			}
		}
	}
}
