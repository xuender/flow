package seq

import (
	"iter"
)

func Chan[E any](input chan E) iter.Seq[E] {
	return func(yield func(E) bool) {
		defer func() {
			_ = recover()
		}()

		for item := range input {
			if !yield(item) {
				close(input)

				return
			}
		}
	}
}
