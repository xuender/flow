package flow_test

import (
	"fmt"

	"gitee.com/xuender/flow"
	"gitee.com/xuender/flow/seq"
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
