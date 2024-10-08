package flow_test

import (
	"fmt"

	"github.com/xuender/flow"
	"github.com/xuender/flow/seq"
)

func ExampleChain() {
	for num := range flow.Chain(
		seq.Range(100),
		flow.Limit[int](10),
		flow.Filter(func(num int) bool { return num%3 == 0 }),
		flow.Map(func(num int) int { return num * 2 }),
	) {
		fmt.Println(num)
	}

	// Output:
	// 0
	// 6
	// 12
	// 18
}

func ExampleChain_mapChangeType() {
	items := flow.Chain(
		seq.Range(100),
		flow.Limit[int](10),
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

func ExampleChain2() {
	for key, val := range flow.Chain2(
		seq.Range2(100),
		flow.Filter2(func(key, _ int) bool { return key%2 == 0 }),
		flow.Limit2[int, int](2),
	) {
		fmt.Println(key, val)
	}

	// Output:
	// 0 0
	// 2 2
}
