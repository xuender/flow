package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleReverse() {
	for num := range seq.Reverse((seq.Range(5))) {
		if num < 3 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 4
	// 3
}
