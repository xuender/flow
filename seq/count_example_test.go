package seq_test

import (
	"fmt"

	"gitee.com/xuender/flow/seq"
)

func ExampleCount() {
	fmt.Println(seq.Count(seq.Filter(seq.Range(6), func(num int) bool { return num%2 == 0 })))

	// Output:
	// 3
}
