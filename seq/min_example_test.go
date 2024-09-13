package seq_test

import (
	"fmt"

	"gitee.com/xuender/flow/seq"
)

func ExampleMin() {
	fmt.Println(seq.Min(seq.Skip[int](seq.Range(6), 3)))

	// Output:
	// 3 true
}
