package seq

import (
	"iter"

	"github.com/xuender/flow/stats"
)

// MovingAvg calculates the moving average of a sequence.
//
// It computes the average of every `size` consecutive elements in the input sequence.
//
// Note:
//
//	Boundary effect at start.
func MovingAvg[V stats.Number](input iter.Seq[V], size int) iter.Seq[V] {
	return func(yield func(V) bool) {
		for items := range Window(input, size) {
			if !yield(stats.Sum(items) / V(len(items))) {
				return
			}
		}
	}
}

// CenteredMovingAvg calculates the centered moving average of a sequence.
//
// It computes the average of `size` consecutive elements, centered when possible.
// The first few and last few averages are computed using available data.
//
// Note:
//
//	Boundary effect at end.
func CenteredMovingAvg[V stats.Number](input iter.Seq[V], size int) iter.Seq[V] {
	return func(yield func(V) bool) {
		cache := []V{}
		pass := 0

		for num := range input {
			cache = append(cache, num)
			length := len(cache)

			switch {
			case length < size:
				pass++

				continue
			case length > size:
				cache = cache[1:]
			}

			if !yield(stats.Sum(cache) / V(len(cache))) {
				return
			}
		}

		for idx := range pass {
			if !yield(stats.Sum(cache[idx:]) / V(len(cache)-idx)) {
				return
			}
		}
	}
}
