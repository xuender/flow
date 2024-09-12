package flow_test

import (
	"fmt"

	"gitee.com/xuender/flow"
)

func ExamplePeek() {
	flow.Emit(flow.Chain(flow.Range(3), flow.Peek(func(num int) {
		fmt.Println(num)
	})))

	// Output:
	// 0
	// 1
	// 2
}
