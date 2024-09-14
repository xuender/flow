package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleCollect() {
	fmt.Println(seq.Collect(seq.Range(3), 3))

	// Output:
	// [0 1 2]
}

func ExampleCollect2() {
	fmt.Println(seq.Collect2(seq.Range2(3), 3))

	// Output:
	// map[0:0 1:1 2:2]
}
