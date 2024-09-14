package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleLimit() {
	for num := range seq.Limit(seq.Range(10), 3) {
		if num > 1 {
			break
		}

		fmt.Println(num)
	}

	for num := range seq.Limit(seq.Range(10), 3) {
		fmt.Println(num)
	}

	// Output:
	// 0
	// 1
	// 0
	// 1
	// 2
}
