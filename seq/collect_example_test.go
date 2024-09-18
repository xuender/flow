package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleCollect() {
	fmt.Println(seq.Collect(seq.Range(4, 7), 3))

	// Output:
	// [4 5 6]
}

func ExampleCollect2() {
	fmt.Println(seq.Collect2(seq.Range2(4, 7), 3))

	// Output:
	// map[0:4 1:5 2:6]
}
