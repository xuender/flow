package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleSkip() {
	for num := range seq.Skip(seq.Range(100), 8) {
		if num > 9 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 8
	// 9
}

func ExampleSkip2() {
	for key, val := range seq.Skip2(seq.Range2(100), 8) {
		if key > 9 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 8 8
	// 9 9
}
