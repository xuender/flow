package seq

import (
	"iter"
	"time"
)

const _defaultDuration = time.Duration(50) * time.Millisecond

// Chan converts a channel into an iterator sequence.
//
// This function takes a channel `input` and returns an iterator function that yields elements
// from the channel until it is closed.
//
// Args:
//
//	input (chan E): The input channel of elements.
//
// Returns:
//
//	iter.Seq[E]: An iterator function yielding elements from the channel.
func Chan[E any](input chan E) iter.Seq[E] {
	return func(yield func(E) bool) {
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

// ToChans distributes elements from the input sequence to multiple channels.
//
// This function takes a sequence `input` and an integer `size` indicating the number of channels.
// It distributes the elements evenly among the specified number of channels.
//
// Args:
//
//	input (iter.Seq[E]): The input sequence of elements.
//	size (int): The number of channels to create.
//
// Returns:
//
//	[]chan E: A slice of channels containing the distributed elements.
func ToChans[E any](input iter.Seq[E], size int) []chan E {
	chans := make([]chan E, size)
	for idx := range size {
		chans[idx] = make(chan E)
	}

	go chanRun(input, chans)

	return chans
}

func chanRun[E any](input iter.Seq[E], chans []chan E) {
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

func allClose(isClose []bool) bool {
	for _, clo := range isClose {
		if !clo {
			return false
		}
	}

	return true
}

func chanSend[E any](item E, chans []chan E, isClose []bool) bool {
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
