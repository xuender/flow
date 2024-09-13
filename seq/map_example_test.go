package seq_test

import (
	"fmt"

	"gitee.com/xuender/flow/seq"
)

func ExampleMap() {
	int2float64 := seq.Map(seq.Range(6), func(num int) float64 { return float64(num * 2) })

	fmt.Println(seq.Max(int2float64))

	// Output:
	// 10 true
}