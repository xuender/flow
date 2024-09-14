package flow

import (
	"iter"
	"sync"

	"github.com/xuender/flow/seq"
)

// Parallel processes the input sequence using multiple workers and applies a series of steps.
//
// This function takes an input sequence `input` and a variadic list of `steps`.
// Each step is applied sequentially to the input sequence.
//
// Parameters:
//
//	numWorkers (int): The number of worker goroutines to use for processing.
//	input (iter.Seq[E]): The input sequence of elements.
//	steps (...Step[E]): A list of transformation steps to apply.
//
// Returns:
//
//	iter.Seq[E]: A new sequence containing the processed results.
func Parallel[E any](numWorkers int, input iter.Seq[E], steps ...Step[E]) iter.Seq[E] {
	chans := seq.ToChans(input, numWorkers)
	output := make(chan E)

	go run(chans, output, steps)

	return seq.Chan(output)
}

func run[E any](chans []chan E, output chan<- E, steps []Step[E]) {
	wait := sync.WaitGroup{}
	wait.Add(len(chans))

	for _, cha := range chans {
		go func(input chan E) {
			defer wait.Done()

			items := seq.Chan(input)

			for _, step := range steps {
				items = step(items)
			}

			for item := range items {
				output <- item
			}
		}(cha)
	}

	wait.Wait()
	close(output)
}
