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

func ExampleLimit2() {
	for key, val := range seq.Limit2(seq.Range2(10), 3) {
		if key > 1 {
			break
		}

		fmt.Println(key, val)
	}

	for key, val := range seq.Limit2(seq.Range2(10), 3) {
		fmt.Println(key, val)
	}

	// Output:
	// 0 0
	// 1 1
	// 0 0
	// 1 1
	// 2 2
}
