package flow_test

import (
	"fmt"
	"slices"

	"gitee.com/xuender/flow"
)

func ExampleFlatMap() {
	for num := range flow.FlatMap[[]int]()(slices.Values([][]int{{1, 2}, {3, 4}})) {
		fmt.Println(num)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
}
