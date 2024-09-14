package seq

import (
	"cmp"
	"iter"
)

func Sum[E cmp.Ordered](input iter.Seq[E]) E {
	var sum E
	for item := range input {
		sum += item
	}

	return sum
}
