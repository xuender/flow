package seq_test

import (
	"fmt"

	"gitee.com/xuender/flow/seq"
)

func ExampleMax() {
	fmt.Println(seq.Max(seq.Filter(seq.Range(6), func(num int) bool { return num%2 == 0 })))

	// Output:
	// 4 true
}
