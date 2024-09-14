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

func ExampleAppend() {
	for num := range flow.Chain(
		seq.Range(2),
		flow.Append(9),
	) {
		fmt.Println(num)
	}

	// Output:
	// 0
	// 1
	// 9
}

func ExampleAppend2() {
	for key, val := range flow.Chain2(
		seq.Range2(2),
		flow.Append2(seq.T(7, 8)),
	) {
		fmt.Println(key, val)
	}

	// Output:
	// 0 0
	// 1 1
	// 7 8
}

func ExampleDistinct2() {
	fmt.Println(seq.Count2(flow.Chain2(
		seq.Range2(3),
		flow.Append2(seq.T(2, 7), seq.T(8, 4)),
		flow.Distinct2[int, int](),
	)))

	// Output:
	// 4
}

func ExampleFilter2() {
	fmt.Println(seq.Count2(flow.Chain2(
		seq.Range2(10),
		flow.Filter2(func(key, _ int) bool {
			return key%3 == 0
		}),
	)))

	// Output:
	// 4
}

func ExampleLimit2() {
	fmt.Println(seq.Count2(flow.Chain2(
		seq.Range2(10),
		flow.Limit2[int, int](4),
	)))

	// Output:
	// 4
}

func ExampleMerge() {
	fmt.Println(seq.Count(flow.Chain(
		seq.Range(3),
		flow.Merge(seq.Range(4), seq.Range(2)),
	)))

	// Output:
	// 9
}

func ExampleMerge2() {
	fmt.Println(seq.Count2(flow.Chain2(
		seq.Range2(3),
		flow.Merge2(seq.Range2(4), seq.Range2(2)),
	)))

	// Output:
	// 9
}

func ExamplePeek2() {
	seq.Emit2(flow.Chain2(
		seq.Range2(3),
		flow.Peek2(func(key, val int) {
			fmt.Println(key, val)
		}),
	))

	// Output:
	// 0 0
	// 1 1
	// 2 2
}

func ExampleReverse2() {
	fmt.Println(seq.First2(flow.Chain2(
		seq.Range2(100),
		flow.Reverse2[int, int](),
	)))

	// Output:
	// 99 99 true
}

func ExampleSkip2() {
	fmt.Println(seq.First2(flow.Chain2(
		seq.Range2(100),
		flow.Skip2[int, int](99),
	)))

	// Output:
	// 99 99 true
}

func ExampleSort2() {
	fmt.Println(seq.First2(flow.Chain2(
		seq.Range2(100),
		flow.Reverse2[int, int](),
		flow.Sort2[int, int](),
	)))

	// Output:
	// 0 0 true
}

func ExampleSortStableFunc() {
	fmt.Println(seq.First(flow.Chain(
		seq.Range(100),
		flow.SortStableFunc(func(num1, num2 int) int {
			return num2 - num1
		}),
	)))

	// Output:
	// 99 true
}

func ExampleSortFunc2() {
	fmt.Println(seq.First2(flow.Chain2(
		seq.Range2(100),
		flow.SortFunc2(func(num1, num2 seq.Tuple[int, int]) int {
			return num2.A - num1.A
		}),
	)))

	// Output:
	// 99 99 true
}

func ExampleSortStableFunc2() {
	fmt.Println(seq.First2(flow.Chain2(
		seq.Range2(100),
		flow.SortStableFunc2(func(num1, num2 seq.Tuple[int, int]) int {
			return num2.A - num1.A
		}),
	)))

	// Output:
	// 99 99 true
}
