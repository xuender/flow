package flow_test

import (
	"fmt"
	"slices"

	"gitee.com/xuender/flow"
)

func ExampleSum() {
	fmt.Println(flow.Sum(flow.Range(101)))
	fmt.Println(flow.Sum(slices.Values([]string{"a", "b", "c"})))

	// Output:
	// 5050
	// abc
}
