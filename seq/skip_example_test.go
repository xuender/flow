package seq_test

import (
	"fmt"

	"gitee.com/xuender/flow/seq"
)

func ExampleSkip() {
	for num := range seq.Skip[int](seq.Range(10), 8) {
		fmt.Println(num)
	}

	// Output:
	// 8
	// 9
}
