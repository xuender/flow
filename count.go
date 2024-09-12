package flow

import (
	"iter"
)

func Count[E any](seq iter.Seq[E]) int {
	count := 0
	for range seq {
		count++
	}

	return count
}
