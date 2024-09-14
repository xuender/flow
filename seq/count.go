package seq

import (
	"iter"
)

func Count[E any](input iter.Seq[E]) int {
	count := 0
	for range input {
		count++
	}

	return count
}
