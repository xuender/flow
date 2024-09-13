package seq_test

import (
	"fmt"

	"gitee.com/xuender/flow/seq"
)

func ExampleFirst() {
	fmt.Println(seq.First(seq.Range(100)))

	// Output:
	// 0 true
}
