package flow_test

import (
	"fmt"

	"gitee.com/xuender/flow"
)

func ExampleMax() {
	filter := flow.Filter(func(num int) bool { return num%2 == 0 })

	fmt.Println(flow.Max(filter(flow.Range(6))))

	// Output:
	// 4 true
}
