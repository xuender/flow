package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleReduce() {
	fmt.Println(seq.Reduce(
		seq.Range(5),
		func(num1, num2 int) int {
			return num1 + num2
		},
	))

	// Output:
	// 10 true
}

func ExampleReduce2() {
	fmt.Println(seq.Reduce2(
		seq.Range2(5),
		func(key1, val1, key2, val2 int) (int, int) {
			return key1 + key2, val1 + val2
		},
	))

	// Output:
	// 10 10 true
}
