package flow

import "iter"

func Chain[E any](input iter.Seq[E], steps ...Step[E]) iter.Seq[E] {
	for _, step := range steps {
		input = step(input)
	}

	return input
}
