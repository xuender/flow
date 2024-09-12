package flow_test

import (
	"fmt"

	"gitee.com/xuender/flow"
)

func ExampleMap() {
	int2float64 := flow.Map(func(num int) float64 { return float64(num * 2) })

	fmt.Println(flow.Max(int2float64(flow.Range(6))))

	// Output:
	// 10 true
}
