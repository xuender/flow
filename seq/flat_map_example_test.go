package seq_test

import (
	"fmt"
	"slices"

	"github.com/xuender/flow/seq"
)

func ExampleFlatMap() {
	for num := range seq.FlatMap((slices.Values([][]int{{1, 2}, {3, 4}}))) {
		fmt.Println(num)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
}
