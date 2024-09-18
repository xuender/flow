package flow_test

import (
	"fmt"

	"github.com/xuender/flow"
	"github.com/xuender/flow/seq"
)

func ExampleChain() {
	for num := range flow.Chain(
		seq.Range(100),
		flow.Limit[int](3),
		flow.Filter(func(num int) bool { return num%2 == 0 }),
	) {
		fmt.Println(num)
	}

	// Output:
	// 0
	// 2
}

func ExampleChain_map() {
	items := flow.Chain(
		seq.Range(100),
		flow.Limit[int](3),
		flow.Filter(func(num int) bool { return num%2 == 0 }),
	)

	for str := range flow.Chain(
		seq.Map(items, func(num int) string { return fmt.Sprintf("num:%d", num) }),
		flow.Limit[string](2),
	) {
		fmt.Println(str)
	}

	// Output:
	// num:0
	// num:2
}
