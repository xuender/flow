package seq_test

import (
	"fmt"
	"slices"

	"gitee.com/xuender/flow/seq"
)

func ExampleReduce() {
	for num := range seq.Reduce((slices.Values([]int{3, 1, 2}))) {
		fmt.Println(num)
	}

	// Output:
	// 3
	// 2
	// 1
}
