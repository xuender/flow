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

func ExampleEach2() {
	seq.Each2(
		seq.Range2(5),
		func(key, val int) bool {
			if key > 2 {
				return false
			}

			fmt.Println(key, val)

			return true
		},
	)

	// Output:
	// 0 0
	// 1 1
	// 2 2
}
