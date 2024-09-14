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
