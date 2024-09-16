package seq_test

import (
	"fmt"
	"maps"
	"slices"

	"github.com/xuender/flow/seq"
)

func ExampleSorted() {
	for num := range seq.Sorted(slices.Values([]int{3, 1, 2, 4})) {
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

func ExampleSorted2() {
	for key, val := range seq.Sorted2(maps.All(map[int]int{1: 1, 2: 2, 3: 3, 4: 4})) {
		if key > 3 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 1 1
	// 2 2
	// 3 3
}

func ExampleSortedFunc() {
	for num := range seq.SortedFunc(
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

func ExampleSortedFunc2() {
	for key, val := range seq.SortedFunc2(
		seq.Range2(10),
		func(item1, item2 seq.Tuple[int, int]) int { return item2.K - item1.K },
	) {
		if key < 7 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 9 9
	// 8 8
	// 7 7
}

func ExampleSortedStableFunc() {
	for num := range seq.SortedStableFunc(
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

func ExampleSortedStableFunc2() {
	for key, val := range seq.SortedStableFunc2(
		seq.Range2(10),
		func(item1, item2 seq.Tuple[int, int]) int { return item2.K - item1.K },
	) {
		if key < 7 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 9 9
	// 8 8
	// 7 7
}
