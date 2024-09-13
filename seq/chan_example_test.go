package seq_test

import (
	"fmt"

	"gitee.com/xuender/flow/seq"
)

func ExampleChan() {
	cha := make(chan int, 3)
	cha <- 1
	cha <- 2
	cha <- 3

	close(cha)

	for item := range seq.Chan(cha) {
		fmt.Println(item)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleToChans() {
	chans := seq.ToChans(seq.Range(10), 3)

	for _, cha := range chans {
		for item := range cha {
			fmt.Println(item)
		}
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
}
