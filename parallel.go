package flow

import (
	"iter"
	"sync"

	"github.com/xuender/flow/seq"
)

func Parallel[E any](size int, items iter.Seq[E], steps ...Step[E]) iter.Seq[E] {
	chans := seq.ToChans(items, size)
	output := make(chan E)

	go parallelRun(chans, output, steps)

	return seq.Chan(output)
}

func parallelRun[E any](chans []chan E, output chan<- E, steps []Step[E]) {
	wait := sync.WaitGroup{}
	wait.Add(len(chans))

	for _, cha := range chans {
		go func(input chan E) {
			items := seq.Chan(input)

			for _, step := range steps {
				items = step(items)
			}

			for item := range items {
				output <- item
			}

			wait.Done()
		}(cha)
	}

	wait.Wait()
	close(output)
}
