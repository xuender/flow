package seq

import (
	"iter"
	"time"
)

const _defaultDuration = time.Duration(50) * time.Millisecond

// Chan function converts a channel 'input' into an iterable sequence of type iter.Seq[E].
// This allows the channel to be used in contexts where a sequence is expected.
//
// Parameters:
//
//	input: A channel of type E from which elements will be read.
//
// Returns:
//
//	An object that implements the iter.Seq[E] interface, allowing iteration over the elements.
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

// ToChans function splits the given sequence 'input' into multiple channels.
// It creates 'size' number of channels and distributes the elements of the sequence among them.
// This function works with elements of any type E.
//
// Parameters:
//
//	input: The sequence to split, of type iter.Seq[E].
//	size: The number of channels to create.
//
// Returns:
//
//	A slice of channels, each receiving elements from the original sequence.
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
