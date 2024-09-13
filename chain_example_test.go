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

func ExampleChain_map() {
	items := flow.Chain(
		seq.Range(100),
		flow.Limit[int](3),
		flow.Filter(func(num int) bool { return num%2 == 0 }),
	)

	fmt.Println(
		seq.Join(flow.Chain(
			seq.Map(items, func(num int) string { return fmt.Sprintf("num:%d", num) }),
			flow.Limit[string](2),
		), ","),
	)

	// Output:
	// num:0,num:2
}
