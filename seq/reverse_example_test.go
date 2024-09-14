package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleReverse() {
	for num := range seq.Reverse(seq.Range(5)) {
		if num < 3 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 4
	// 3
}

func ExampleReverse2() {
	for key, val := range seq.Reverse2(seq.Range2(5)) {
		if key < 3 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 4 4
	// 3 3
}
