package flow_test

import (
	"fmt"
	"slices"

	"gitee.com/xuender/flow"
)

func ExampleUnique() {
	for num := range flow.Unique[int]()(slices.Values([]int{1, 2, 2, 3, 1})) {
		fmt.Println(num)
	}

	// Output:
	// 1
	// 2
	// 3
}
