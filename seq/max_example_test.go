package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleMax() {
	fmt.Println(seq.Max(seq.Filter(
		seq.Range(6),
		func(num int) bool {
			return num%2 == 0
		},
	)))

	// Output:
	// 4 true
}

func ExampleMax2() {
	fmt.Println(seq.Max2(seq.Filter2(
		seq.Range2(6),
		func(key, _ int) bool {
			return key%2 == 0
		},
	)))

	// Output:
	// 4 4 true
}

func ExampleMaxFunc() {
	fmt.Println(seq.MaxFunc(
		seq.Range(6),
		func(item1, item2 int) int { return item2 - item1 },
	))

	// Output:
	// 0 true
}

func ExampleMaxFunc2() {
	fmt.Println(seq.MaxFunc2(
		seq.Range2(6),
		func(item1, item2 seq.Tuple[int, int]) int { return item2.K - item1.K },
	))

	// Output:
	// 0 0 true
}
