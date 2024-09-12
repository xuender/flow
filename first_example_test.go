package flow_test

import (
	"fmt"

	"gitee.com/xuender/flow"
)

func ExampleFirst() {
	fmt.Println(flow.First(flow.Range(100)))

	// Output:
	// 0 true
}
