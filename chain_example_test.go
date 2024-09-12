package flow_test

import (
	"fmt"

	"gitee.com/xuender/flow"
)

func ExampleChain() {
	fmt.Println(flow.Max(flow.Chain(
		flow.Range(6),
		flow.Limit[int](3),
		flow.Filter(func(num int) bool { return num%2 == 0 }),
	)))

	// Output:
	// 2 true
}
