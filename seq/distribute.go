package seq

import (
	"iter"
	"sync"
	"sync/atomic"
)

// Distribute distributes the input sequence across multiple sequences based on the given number.
//
// It returns the resulting slice of sequences.
//
// Play: https://go.dev/play/p/R16UoqYOJmk
func Distribute[V any](input iter.Seq[V], num int) []iter.Seq[V] {
	var (
		count = int32(num) // nolint: gosec
		lock  sync.Mutex
	)

	ret := make([]iter.Seq[V], num)
	next, stop := iter.Pull(input)

	for idx := range num {
		ret[idx] = func(yield func(V) bool) {
			for {
				lock.Lock()
				item, has := next()
				lock.Unlock()

				if !has {
					return
				}

				if !yield(item) {
					atomic.AddInt32(&count, -1)

					if atomic.LoadInt32(&count) < 1 {
						stop()
					}

					return
				}
			}
		}
	}

	return ret
}

// Distribute2 distributes the input sequence across multiple sequences based on the given number.
//
// It returns the resulting slice of sequences.
//
// Play: https://go.dev/play/p/K7MnBXCMwiX
func Distribute2[K, V any](input iter.Seq2[K, V], num int) []iter.Seq2[K, V] {
	var (
		count = int32(num) // nolint: gosec
		lock  sync.Mutex
	)

	ret := make([]iter.Seq2[K, V], num)
	next, stop := iter.Pull2(input)

	for idx := range num {
		ret[idx] = func(yield func(K, V) bool) {
			for {
				lock.Lock()
				key, val, has := next()
				lock.Unlock()

				if !has {
					return
				}

				if !yield(key, val) {
					atomic.AddInt32(&count, -1)

					if atomic.LoadInt32(&count) < 1 {
						stop()
					}

					return
				}
			}
		}
	}

	return ret
}
