package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleEmit() {
	seq.Emit(seq.Peek(seq.Range(3), func(num int) {
		fmt.Println(num)
	}))

	// Output:
	// 0
	// 1
	// 2
}
