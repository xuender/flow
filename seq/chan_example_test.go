package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleChan() {
	cha := make(chan int, 3)
	cha <- 1
	cha <- 2
	cha <- 3

	close(cha)

	for item := range seq.Chan(cha) {
		if item > 1 {
			break
		}

		fmt.Println(item)
	}

	// Output:
	// 1
}

func ExampleToChans() {
	chans := seq.ToChans(seq.Range(10), 3)

	for _, cha := range chans {
		for item := range cha {
			if item > 2 {
				break
			}

			fmt.Println(item)
		}
	}

	// Output:
	// 0
	// 1
	// 2
}
