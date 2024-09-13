package seq

import (
	"cmp"
	"iter"
)

func Max[E cmp.Ordered](seq iter.Seq[E]) (E, bool) {
	var maxItem E

	has := false

	for item := range seq {
		if !has || item > maxItem {
			maxItem = item
			has = true
		}
	}

	return maxItem, has
}
