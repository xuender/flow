package seq_test

import (
	"fmt"
	"slices"

	"github.com/xuender/flow/seq"
)

func ExampleDistinct() {
	for num := range seq.Distinct(slices.Values([]int{1, 2, 2, 3, 1})) {
		if num > 2 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 1
	// 2
}

func ExampleDistinct2() {
	for key, val := range seq.Distinct2(seq.Merge2(
		seq.Range2(3),
		seq.Range2(1, 7),
	)) {
		if key > 2 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 0 0
	// 1 1
	// 2 2
}
