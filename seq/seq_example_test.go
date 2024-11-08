package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleSeq() {
	for val := range seq.Seq(1, 2, 3) {
		fmt.Println(val)
	}

	// Output:
	// 1
	// 2
	// 3
}

func ExampleSeq2() {
	for key, val := range seq.Seq2(seq.T(1, "one"), seq.T(2, "two"), seq.T(3, "three")) {
		if key > 2 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 1 one
	// 2 two
}
