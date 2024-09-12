package flow_test

import (
	"fmt"

	"gitee.com/xuender/flow"
)

func ExampleChan() {
	cha := make(chan int, 3)
	cha <- 1
	cha <- 2
	cha <- 3

	close(cha)

	for item := range flow.Chan(cha) {
		fmt.Println(item)
	}

	// Output:
	// 1
	// 2
	// 3
}
