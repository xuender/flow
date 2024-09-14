package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleEach() {
	seq.Each(
		seq.Range(5),
		func(num int) bool {
			if num > 2 {
				return false
			}

			fmt.Println(num)

			return true
		},
	)

	// Output:
	// 0
	// 1
	// 2
}
