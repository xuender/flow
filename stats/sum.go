package stats

import "cmp"

// Sum calculates the sum of a slice of ordered numeric values.
//
// It iterates over the slice and accumulates the total sum.
func Sum[V cmp.Ordered](items []V) V {
	var ret V

	for _, num := range items {
		ret += num
	}

	return ret
}
