package flow_test

import (
	"fmt"

	"gitee.com/xuender/flow"
)

func ExampleReverse() {
	for num := range flow.Reverse[int]()(flow.Range(5)) {
		fmt.Println(num)
	}

	// Output:
	// 4
	// 3
	// 2
	// 1
	// 0
}
