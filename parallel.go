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
// Args:
//
//	numWorkers int: The number of worker goroutines to use for processing.
//	input iter.Seq[V]: The input sequence of elements.
//	steps ...Step[V]: A list of transformation steps to apply.
//
// Returns:
//
//	iter.Seq[V]: A new sequence containing the processed results.
func Parallel[V any](numWorkers int, input iter.Seq[V], steps ...Step[V]) iter.Seq[V] {
	chans := seq.ToChans(input, numWorkers)
	output := make(chan V)

	go run(chans, output, steps)

	return seq.Chan(output)
}

// Parallel2 processes the input sequence in parallel using multiple workers.
//
// It returns a new sequence with the processed results.
//
// Args:
//
//	numWorkers int: The number of workers.
//	input iter.Seq2[K, V]: The input sequence.
//	steps ...Step2[K, V]: The processing steps.
//
// Returns:
//
//	iter.Seq2[K, V]: A new sequence with results.
func Parallel2[K, V any](numWorkers int, input iter.Seq2[K, V], steps ...Step2[K, V]) iter.Seq2[K, V] {
	chans := seq.ToChans2(input, numWorkers)
	output := make(chan seq.Tuple[K, V])

	go run2(chans, output, steps)

	return seq.Chan2(output)
}

func run[V any](chans []chan V, output chan<- V, steps []Step[V]) {
	wait := sync.WaitGroup{}
	wait.Add(len(chans))

	for _, cha := range chans {
		go func(input chan V) {
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

func run2[K, V any](chans []chan seq.Tuple[K, V], output chan<- seq.Tuple[K, V], steps []Step2[K, V]) {
	wait := sync.WaitGroup{}
	wait.Add(len(chans))

	for _, cha := range chans {
		go func(input chan seq.Tuple[K, V]) {
			defer wait.Done()

			items := seq.Chan2(input)

			for _, step := range steps {
				items = step(items)
			}

			for key, val := range items {
				output <- seq.T(key, val)
			}
		}(cha)
	}

	wait.Wait()
	close(output)
}
