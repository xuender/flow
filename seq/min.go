package seq

import (
	"cmp"
	"iter"
)

func Min[E cmp.Ordered](seq iter.Seq[E]) (E, bool) {
	var minItem E

	has := false

	for item := range seq {
		if !has || item < minItem {
			minItem = item
			has = true
		}
	}

	return minItem, has
}
