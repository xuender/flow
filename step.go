package flow

import "iter"

type Step[I, O any] func(iter.Seq[I]) iter.Seq[O]
