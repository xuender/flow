package seq_test

import (
	"fmt"

	"gitee.com/xuender/flow/seq"
)

func ExampleReverse() {
	for num := range seq.Reverse[int]((seq.Range(5))) {
		fmt.Println(num)
	}

	// Output:
	// 4
	// 3
	// 2
	// 1
	// 0
}
