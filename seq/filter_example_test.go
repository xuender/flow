package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleFilter() {
	for num := range seq.Filter(
		seq.Range(6),
		func(num int) bool { return num%2 == 0 },
	) {
		if num > 2 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 0
	// 2
}

func ExampleFilter2() {
	for key, val := range seq.Filter2(
		seq.Range2(6),
		func(key, _ int) bool { return key%2 == 0 },
	) {
		if key > 2 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 0 0
	// 2 2
}
