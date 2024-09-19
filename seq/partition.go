package seq

import (
	"iter"
	"sync"
)

// Partition divides the input sequence into two sequences based on the predicate.
//
// It returns two iter.Seq[V] the matched and unmatched sequences.
func Partition[V any](input iter.Seq[V], predicate func(V) bool) (iter.Seq[V], iter.Seq[V]) {
	var (
		matchedYield, unmatchedYield func(V) bool
		wait, ready                  sync.WaitGroup
		readyNum                     = 2
	)

	wait.Add(1)
	ready.Add(readyNum)

	go func() {
		ready.Wait()

		defer wait.Done()

		for item := range input {
			if predicate(item) {
				if !matchedYield(item) {
					return
				}
			} else {
				if !unmatchedYield(item) {
					return
				}
			}
		}
	}()

	return func(yield func(V) bool) {
			matchedYield = yield

			ready.Done()
			wait.Wait()
		}, func(yield func(V) bool) {
			unmatchedYield = yield

			ready.Done()
			wait.Wait()
		}
}

// Partition divides the input sequence into two sequences based on the predicate.
//
// Returns two iter.Seq2[K, V] matched and unmatched sequences.
func Partition2[K, V any](input iter.Seq2[K, V], predicate func(K, V) bool) (iter.Seq2[K, V], iter.Seq2[K, V]) {
	var (
		matchedYield, unmatchedYield func(K, V) bool
		wait, ready                  sync.WaitGroup
		readyNum                     = 2
	)

	wait.Add(1)
	ready.Add(readyNum)

	go func() {
		ready.Wait()

		defer wait.Done()

		for key, val := range input {
			if predicate(key, val) {
				if !matchedYield(key, val) {
					return
				}
			} else {
				if !unmatchedYield(key, val) {
					return
				}
			}
		}
	}()

	return func(yield func(K, V) bool) {
			matchedYield = yield

			ready.Done()
			wait.Wait()
		}, func(yield func(K, V) bool) {
			unmatchedYield = yield

			ready.Done()
			wait.Wait()
		}
}
