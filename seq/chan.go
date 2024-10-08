package seq

import (
	"iter"
	"time"
)

const _defaultDuration = time.Duration(50) * time.Millisecond

// Chan converts a channel into an iterator sequence.
//
// It takes a channel `input` and returns an iterator function that yields elements
// from the channel until it is closed.
//
// Play: https://go.dev/play/p/Mi1zPueGIj5
func Chan[V any](input chan V) iter.Seq[V] {
	return func(yield func(V) bool) {
		defer func() {
			_ = recover()
		}()

		for item := range input {
			if !yield(item) {
				close(input)

				return
			}
		}
	}
}

// Chan2 converts a channel of (key, value) tuples into a sequence.
//
// It returns a function that yields the tuples.
//
// Play: https://go.dev/play/p/Rut7y-2vab8
func Chan2[K, V any](input chan Tuple[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		defer func() {
			_ = recover()
		}()

		for item := range input {
			if !yield(item.K, item.V) {
				close(input)

				return
			}
		}
	}
}

// ToChans distributes elements from the input sequence to multiple channels.
//
// It takes a sequence `input` and an integer `size` indicating the number of channels.
// It distributes the elements evenly among the specified number of channels.
//
// Play: https://go.dev/play/p/jGnmFmeCwgj
func ToChans[V any](input iter.Seq[V], size int) []chan V {
	chans := make([]chan V, size)
	for idx := range size {
		chans[idx] = make(chan V)
	}

	go chanRun(input, chans)

	return chans
}

// ToChans2 converts the input sequence into a slice of channels.
//
// Each channel receives (key, value) tuples.
//
// Play: https://go.dev/play/p/3BPhqH9_H3H
func ToChans2[K, V any](input iter.Seq2[K, V], size int) []chan Tuple[K, V] {
	chans := make([]chan Tuple[K, V], size)
	for idx := range size {
		chans[idx] = make(chan Tuple[K, V])
	}

	go chanRun2(input, chans)

	return chans
}

func chanRun[V any](input iter.Seq[V], chans []chan V) {
	isClose := make([]bool, len(chans))

	for item := range input {
		if chanSend(item, chans, isClose) {
			break
		}
	}

	for idx, cha := range chans {
		if isClose[idx] {
			continue
		}

		close(cha)
	}
}

func chanRun2[K, V any](input iter.Seq2[K, V], chans []chan Tuple[K, V]) {
	isClose := make([]bool, len(chans))

	for key, val := range input {
		if chanSend(T(key, val), chans, isClose) {
			break
		}
	}

	for idx, cha := range chans {
		if isClose[idx] {
			continue
		}

		close(cha)
	}
}

func allClose(isClose []bool) bool {
	for _, clo := range isClose {
		if !clo {
			return false
		}
	}

	return true
}

func chanSend[V any](item V, chans []chan V, isClose []bool) bool {
	var closeIdx int

	defer func() {
		if err := recover(); err != nil {
			isClose[closeIdx] = true
		}
	}()

	for {
		if allClose(isClose) {
			return true
		}

		for idx, cha := range chans {
			if isClose[idx] {
				continue
			}

			closeIdx = idx

			select {
			case cha <- item:
				return false
			default:
				continue
			}
		}

		time.Sleep(_defaultDuration)
	}
}
