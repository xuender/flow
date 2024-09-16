package flow_test

import (
	"fmt"

	"github.com/xuender/flow"
	"github.com/xuender/flow/seq"
)

func ExampleParallel() {
	for num := range flow.Parallel(2,
		seq.Range(100),
		flow.Limit[int](8),
		flow.Filter(func(num int) bool { return num%3 == 0 }),
		flow.Sort[int](),
	) {
		if num > 3 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 0
	// 3
}

func ExampleParallel2() {
	for key, val := range flow.Parallel2(2,
		seq.Range2(100),
		flow.Limit2[int, int](8),
		flow.Filter2(func(key, _ int) bool { return key%3 == 0 }),
		flow.Sort2[int, int](),
	) {
		if key > 3 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 0 0
	// 3 3
}
