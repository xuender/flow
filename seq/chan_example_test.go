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

func ExampleChan2() {
	cha := make(chan seq.Tuple[int, int], 3)
	cha <- seq.T(1, 1)
	cha <- seq.T(2, 2)
	cha <- seq.T(3, 3)

	close(cha)

	for key, val := range seq.Chan2(cha) {
		if key > 1 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 1 1
}

func ExampleToChans() {
	chans := seq.ToChans(seq.Range(5), 3)

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
}

func ExampleToChans2() {
	chans := seq.ToChans2(seq.Range2(5), 3)

	for _, cha := range chans {
		for item := range cha {
			fmt.Println(item.K, item.V)
		}
	}

	// Output:
	// 0 0
	// 1 1
	// 2 2
	// 3 3
	// 4 4
}
