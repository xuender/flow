package flow_test

import (
	"fmt"

	"github.com/xuender/flow"
	"github.com/xuender/flow/seq"
)

func ExampleParallel() {
	for num := range flow.Parallel(
		2,
		seq.Range(100),
		flow.Limit[int](3),
		flow.Filter(func(num int) bool { return num%3 == 0 }),
	) {
		fmt.Println(num)
	}

	// Output:
	// 0
	// 3
}

func ExampleParallel2() {
	for key, val := range flow.Parallel2(
		2,
		seq.Range2(100),
		flow.Limit2[int, int](3),
		flow.Filter2(func(key, _ int) bool { return key%3 == 0 }),
	) {
		fmt.Println(key, val)
	}

	// Output:
	// 0 0
	// 3 3
}
