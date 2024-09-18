package flow

import (
	"context"
	"iter"
	"sync"

	"github.com/xuender/flow/seq"
)

// Parallel processes the input sequence using multiple workers and applies a series of steps.
//
// It returns a new sequence with the processed results.
func Parallel[V any](num int, input iter.Seq[V], steps ...Step[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		group := []Step[V]{}
		isParallel := false

		for _, step := range steps {
			if len(group) == 0 || isParallel == step.Parallel {
				group = append(group, step)
			} else {
				if isParallel {
					input = parallel(num, input, group)
				} else {
					input = Chain(input, group...)
				}

				group = []Step[V]{step}
			}

			isParallel = step.Parallel
		}

		if isParallel {
			input = parallel(num, input, group)
		} else {
			input = Chain(input, group...)
		}

		for item := range input {
			if !yield(item) {
				return
			}
		}
	}
}

func parallel[V any](num int, input iter.Seq[V], steps []Step[V]) iter.Seq[V] {
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
func Parallel2[K, V any](num int, input iter.Seq2[K, V], steps ...Step2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		group := []Step2[K, V]{}
		isParallel := false

		for _, step := range steps {
			if len(group) == 0 || isParallel == step.Parallel {
				group = append(group, step)
			} else {
				if isParallel {
					input = parallel2(num, input, group)
				} else {
					input = Chain2(input, group...)
				}

				group = []Step2[K, V]{step}
			}

			isParallel = step.Parallel
		}

		if isParallel {
			input = parallel2(num, input, group)
		} else {
			input = Chain2(input, group...)
		}

		for key, val := range input {
			if !yield(key, val) {
				return
			}
		}
	}
}

func parallel2[K, V any](num int, input iter.Seq2[K, V], steps []Step2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seqs := seq.Distribute2(input, num)
		output := make(chan seq.Tuple[K, V], num)
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
