package flow

import "iter"

func FlatMap[S ~[]E, E any]() Step[S, E] {
	return func(input iter.Seq[S]) iter.Seq[E] {
		return func(yield func(E) bool) {
			for slice := range input {
				for _, item := range slice {
					if !yield(item) {
						return
					}
				}
			}
		}
	}
}
