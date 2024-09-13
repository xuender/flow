package flow

import (
	"iter"
	"time"
)

const _defaultDuration = time.Duration(50) * time.Millisecond

func Chans[E any](input iter.Seq[E], size int) []chan E {
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

	for idx, ch := range chans {
		if isClose[idx] {
			continue
		}

		close(ch)
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
