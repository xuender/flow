package seq_test

import (
	"fmt"
	"slices"

	"github.com/xuender/flow/seq"
)

func ExampleDistinct() {
	for num := range seq.Distinct((slices.Values([]int{1, 2, 2, 3, 1}))) {
		fmt.Println(num)
	}

	// Output:
	// 1
	// 2
	// 3
}
