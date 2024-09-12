package flow_test

import (
	"fmt"
	"slices"

	"gitee.com/xuender/flow"
)

func ExampleSort() {
	for num := range flow.Sort[int]()(slices.Values([]int{3, 1, 2})) {
		fmt.Println(num)
	}

	// Output:
	// 1
	// 2
	// 3
}
