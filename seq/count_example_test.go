package seq_test

import (
	"fmt"

	"gitee.com/xuender/flow/seq"
)

func ExampleCount() {
	fmt.Println(seq.Count(seq.Range(6)))

	// Output:
	// 6
}
