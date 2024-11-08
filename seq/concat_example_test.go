package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleConcat() {
	for val := range seq.Concat(seq.Seq(1, 2), seq.Seq(5, 6)) {
		if val > 5 {
			break
		}

		fmt.Println(val)
	}

	// Output:
	// 1
	// 2
	// 5
}

func ExampleConcat2() {
	for key, val := range seq.Concat2(
		seq.Seq2(seq.T(1, "one"), seq.T(2, "two")),
		seq.Seq2(seq.T(5, "five"), seq.T(6, "six")),
	) {
		if key > 5 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 1 one
	// 2 two
	// 5 five
}
