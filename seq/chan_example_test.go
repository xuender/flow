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
