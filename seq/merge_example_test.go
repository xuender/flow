package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleMerge() {
	for num := range seq.Merge(seq.Range(3), seq.Range(6, 10)) {
		if num > 6 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 0
	// 1
	// 2
	// 6
}

func ExampleMerge2() {
	for key, val := range seq.Merge2(seq.Range2(3), seq.Range2(6, 10)) {
		if val > 6 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 0 0
	// 1 1
	// 2 2
	// 0 6
}
