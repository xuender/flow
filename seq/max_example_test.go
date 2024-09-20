package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleMax() {
	fmt.Println(seq.Max(seq.Filter(
		seq.Range(6),
		func(num int) bool {
			return num%2 == 0
		},
	)))

	// Output:
	// 4 true
}

func ExampleMax2() {
	fmt.Println(seq.Max2(seq.Filter2(
		seq.Range2(6),
		func(key, _ int) bool {
			return key%2 == 0
		},
	)))

	// Output:
	// 4 4 true
}
