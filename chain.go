package flow

import "iter"

func Chain[E any](seq iter.Seq[E], steps ...Step[E, E]) iter.Seq[E] {
	for _, step := range steps {
		seq = step(seq)
	}

	return seq
}
