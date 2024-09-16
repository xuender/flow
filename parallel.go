package flow

import (
	"context"
	"iter"
	"sync"

	"github.com/xuender/flow/seq"
)

// Parallel processes the input sequence using multiple workers and applies a series of steps.
//
// This function takes an input sequence `input` and a variadic list of `steps`.
// Each step is applied sequentially to the input sequence.
//
// Note:
//
//	The func has a defect; use with caution.
//	I need help!
func Parallel[V any](num int, input iter.Seq[V], steps ...Step[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		seqs := seq.Distribute(input, num)
		output := make(chan V, num)
		wait := &sync.WaitGroup{}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		wait.Add(num)

		go closeChan(ctx, wait, output)

		for idx := range num {
			go func() {
				select {
				case <-ctx.Done():
				default:
					seq.Send(Chain(seqs[idx], steps...), output)
					wait.Done()
				}
			}()
		}

		for item := range output {
			if !yield(item) {
				return
			}
		}
	}
}

func closeChan[V any](ctx context.Context, wait *sync.WaitGroup, cha chan V) {
	select {
	case <-ctx.Done():
	default:
		wait.Wait()
	}

	close(cha)
}

// Parallel2 processes the input sequence in parallel using multiple workers.
//
// It returns a new sequence with the processed results.
//
// Note:
//
//	The func has a defect; use with caution.
//	I need help!
func Parallel2[K, V any](num int, input iter.Seq2[K, V], steps ...Step2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seqs := seq.Distribute2(input, num)
		output := make(chan seq.Tuple[K, V])
		wait := &sync.WaitGroup{}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		wait.Add(num)

		go closeChan(ctx, wait, output)

		for idx := range num {
			go func() {
				select {
				case <-ctx.Done():
				default:
					seq.Send2(Chain2(seqs[idx], steps...), output)
					wait.Done()
				}
			}()
		}

		for item := range output {
			if !yield(item.K, item.V) {
				return
			}
		}
	}
}
