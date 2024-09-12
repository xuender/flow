package flow_test

import (
	"fmt"

	"gitee.com/xuender/flow"
)

func ExampleFilter() {
	filter := flow.Filter(func(i int) bool { return i%2 == 0 })

	fmt.Println(flow.Max(filter(flow.Range(6))))

	// Output:
	// 4 true
}
