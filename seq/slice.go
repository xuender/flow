package seq

import (
	"iter"
)

func Slice[E any](input iter.Seq[E], size int) []E {
	ret := make([]E, 0, size)

	for item := range input {
		ret = append(ret, item)
	}

	return ret
}
