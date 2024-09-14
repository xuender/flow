package seq_test

import (
	"fmt"
	"slices"

	"github.com/xuender/flow/seq"
)

func ExampleSort() {
	for num := range seq.Sort((slices.Values([]int{3, 1, 2, 4}))) {
		if num > 3 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleSortFunc() {
	for num := range seq.SortFunc(
		seq.Range(10),
		func(item1, item2 int) int { return item2 - item1 },
	) {
		if num < 7 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 9
	// 8
	// 7
}
