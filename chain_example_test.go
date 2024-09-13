package flow_test

import (
	"fmt"

	"gitee.com/xuender/flow"
	"gitee.com/xuender/flow/seq"
)

func ExampleChain() {
	fmt.Println(seq.Max(flow.Chain(
		seq.Range(100),
		flow.Limit[int](3),
		flow.Filter(func(num int) bool { return num%2 == 0 }),
	)))

	// Output:
	// 2 true
}
