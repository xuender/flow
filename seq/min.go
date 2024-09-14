package seq

import (
	"cmp"
	"iter"
)

func Min[E cmp.Ordered](input iter.Seq[E]) (E, bool) {
	var minItem E

	has := false

	for item := range input {
		if !has || item < minItem {
			minItem = item
			has = true
		}
	}

	return minItem, has
}
