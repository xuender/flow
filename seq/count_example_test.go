package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleCount() {
	fmt.Println(seq.Count(seq.Range(6)))

	// Output:
	// 6
}

func ExampleCount2() {
	fmt.Println(seq.Count2(seq.Range2(6)))

	// Output:
	// 6
}
