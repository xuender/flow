package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleEmit() {
	seq.Emit(seq.Peek(
		seq.Range(3),
		func(num int) {
			fmt.Println(num)
		},
	))

	// Output:
	// 0
	// 1
	// 2
}

func ExampleEmit2() {
	seq.Emit2(seq.Peek2(
		seq.Range2(3),
		func(key, val int) {
			fmt.Println(key, val)
		},
	))

	// Output:
	// 0 0
	// 1 1
	// 2 2
}
