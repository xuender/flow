package flow_test

import (
	"fmt"

	"gitee.com/xuender/flow"
)

func ExampleCount() {
	filter := flow.Filter(func(num int) bool { return num%2 == 0 })

	fmt.Println(flow.Count(filter(flow.Range(6))))

	// Output:
	// 3
}
