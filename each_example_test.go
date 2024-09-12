package flow_test

import (
	"fmt"

	"gitee.com/xuender/flow"
)

func ExampleEach() {
	flow.Each(flow.Range(5), func(num int) bool {
		fmt.Println(num)

		return true
	})

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}
