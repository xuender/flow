package seq_test

import (
	"fmt"
	"slices"

	"gitee.com/xuender/flow/seq"
)

func ExampleSort() {
	for num := range seq.Sort((slices.Values([]int{3, 1, 2}))) {
		fmt.Println(num)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleSortFunc() {
	for num := range seq.SortFunc(
		seq.Range(3),
		func(item1, item2 int) int { return item2 - item1 },
	) {
		fmt.Println(num)
	}

	// Output:
	// 2
	// 1
	// 0
}
