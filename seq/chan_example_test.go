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

func ExampleToChans2() {
	chans := seq.ToChans2(seq.Range2(10), 3)

	for _, cha := range chans {
		for item := range cha {
			if item.A > 2 {
				break
			}

			fmt.Println(item.A, item.B)
		}
	}

	// Output:
	// 0 0
	// 1 1
	// 2 2
}
