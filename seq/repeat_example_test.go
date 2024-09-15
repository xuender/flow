package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleRepeat() {
	for num := range seq.Repeat(seq.Range(3), 2) {
		if num > 1 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 0
	// 0
	// 1
	// 1
}

func ExampleRepeat2() {
	for key, val := range seq.Repeat2(seq.Range2(7, 10), 2) {
		if key > 1 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 0 7
	// 0 7
	// 1 8
	// 1 8
}
