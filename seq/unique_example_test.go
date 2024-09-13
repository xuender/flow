package seq_test

import (
	"fmt"
	"slices"

	"github.com/xuender/flow/seq"
)

func ExampleUnique() {
	for num := range seq.Unique[int]((slices.Values([]int{1, 2, 2, 3, 1}))) {
		fmt.Println(num)
	}

	// Output:
	// 1
	// 2
	// 3
}
