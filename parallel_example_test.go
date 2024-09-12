package flow_test

import (
	"fmt"

	"gitee.com/xuender/flow"
)

func ExampleParallel() {
	for num := range flow.Parallel(
		2,
		flow.Range(10),
		flow.Filter(func(num int) bool { return num%3 == 0 }),
	) {
		fmt.Println(num)
	}

	// Output:
	// 0
	// 3
	// 6
	// 9
}
