package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleMerge() {
	for num := range seq.Merge(seq.Range(3), seq.Range(6, 8)) {
		fmt.Println(num)
	}

	// Output:
	// 0
	// 1
	// 2
	// 6
	// 7
}
