package seq_test

import (
	"fmt"

	"gitee.com/xuender/flow/seq"
)

func ExampleLimit() {
	fmt.Println(seq.Join(seq.Limit(seq.Range(10), 3), ","))

	// Output:
	// 0,1,2
}
