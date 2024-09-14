package seq_test

import (
	"fmt"
	"maps"
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
	for key, num := range seq.Distinct2(seq.Sort2(maps.All(map[int]int{1: 11, 2: 22, 3: 33}))) {
		if key > 2 {
			break
		}

		fmt.Println(key, num)
	}

	// Output:
	// 1 11
	// 2 22
}
