package flow_test

import (
	"fmt"
	"slices"

	"github.com/xuender/flow"
	"github.com/xuender/flow/seq"
)

func ExampleDistinct() {
	fmt.Println(seq.Sum(flow.Chain(
		slices.Values([]int{1, 2, 2, 3, 3}),
		flow.Distinct[int](),
	)))

	// Output:
	// 6
}

func ExampleFilter() {
	fmt.Println(seq.Sum(flow.Chain(
		seq.Range(10),
		flow.Filter(func(i int) bool {
			return i%3 == 0
		}),
	)))

	// Output:
	// 18
}

func ExampleLimit() {
	fmt.Println(seq.Sum(flow.Chain(
		seq.Range(10),
		flow.Limit[int](4),
	)))

	// Output:
	// 6
}

func ExamplePeek() {
	seq.Emit(flow.Chain(
		seq.Range(3),
		flow.Peek(func(num int) {
			fmt.Println(num)
		}),
	))

	// Output:
	// 0
	// 1
	// 2
}

func ExampleReverse() {
	fmt.Println(seq.First(flow.Chain(
		seq.Range(100),
		flow.Reverse[int](),
	)))

	// Output:
	// 99 true
}

func ExampleSkip() {
	fmt.Println(seq.First(flow.Chain(
		seq.Range(100),
		flow.Skip[int](99),
	)))

	// Output:
	// 99 true
}

func ExampleSort() {
	fmt.Println(seq.First(flow.Chain(
		seq.Range(100),
		flow.Reverse[int](),
		flow.Sort[int](),
	)))

	// Output:
	// 0 true
}

func ExampleSortFunc() {
	fmt.Println(seq.First(flow.Chain(
		seq.Range(100),
		flow.SortFunc(func(num1, num2 int) int {
			return num2 - num1
		}),
	)))

	// Output:
	// 99 true
}
