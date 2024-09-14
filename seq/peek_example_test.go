package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExamplePeek() {
	for num := range seq.Peek(seq.Range(3), func(num int) {
		fmt.Println(num)
	}) {
		if num > 0 {
			break
		}
	}

	// Output:
	// 0
	// 1
}

func ExamplePeek2() {
	for num := range seq.Peek2(seq.Range2(3), func(key, val int) {
		fmt.Println(key, val)
	}) {
		if num > 0 {
			break
		}
	}

	// Output:
	// 0 0
	// 1 1
}
