package flow_test

import (
	"fmt"

	"gitee.com/xuender/flow"
)

func ExampleSkip() {
	for num := range flow.Chain(flow.Range(10), flow.Skip[int](8)) {
		fmt.Println(num)
	}

	// Output:
	// 8
	// 9
}
