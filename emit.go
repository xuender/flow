package flow

import "iter"

func Emit[E any](seq iter.Seq[E], steps ...Step[E, E]) {
	for _, step := range steps {
		seq = step(seq)
	}

	call := func(E) bool {
		return true
	}

	seq(call)
}
