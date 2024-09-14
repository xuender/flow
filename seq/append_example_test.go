package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleAppend() {
	for num := range seq.Append(seq.Range(1, 3), 4, 5) {
		fmt.Println(num)
	}

	// Output:
	// 1
	// 2
	// 4
	// 5
}
