package seq_test

import (
	"fmt"
	"slices"

	"gitee.com/xuender/flow/seq"
)

func ExampleFlatMap() {
	for num := range seq.FlatMap[[]int]((slices.Values([][]int{{1, 2}, {3, 4}}))) {
		fmt.Println(num)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
}
