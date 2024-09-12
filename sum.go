package flow

import "iter"

func Sum[E int](seq iter.Seq[E]) E {
	var sum E
	for item := range seq {
		sum += item
	}

	return sum
}
