package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleFilter() {
	for num := range seq.Filter(
		seq.Range(6),
		func(num int) bool {
			return num%2 == 0
		},
	) {
		if num > 2 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 0
	// 2
}
