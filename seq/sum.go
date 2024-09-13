package seq

import (
	"cmp"
	"iter"
)

func Sum[E cmp.Ordered](seq iter.Seq[E]) E {
	var sum E
	for item := range seq {
		sum += item
	}

	return sum
}
