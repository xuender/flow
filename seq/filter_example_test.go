package seq_test

import (
	"fmt"

	"gitee.com/xuender/flow/seq"
)

func ExampleFilter() {
	fmt.Println(seq.Max(seq.Filter(seq.Range(6), func(i int) bool { return i%2 == 0 })))

	// Output:
	// 4 true
}