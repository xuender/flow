package seq_test

import (
	"fmt"
	"slices"

	"gitee.com/xuender/flow/seq"
)

func ExampleSort() {
	for num := range seq.Sort[int]((slices.Values([]int{3, 1, 2}))) {
		fmt.Println(num)
	}

	// Output:
	// 1
	// 2
	// 3
}
