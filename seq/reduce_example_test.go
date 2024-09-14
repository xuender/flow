package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleReduce() {
	fmt.Println(seq.Reduce(
		seq.Range(5),
		func(num1, num2 int) int {
			return num1 + num2
		},
	))

	// Output:
	// 10 true
}
