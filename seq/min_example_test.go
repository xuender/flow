package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleMin() {
	fmt.Println(seq.Min(seq.Skip(seq.Range(6), 3)))

	// Output:
	// 3 true
}

func ExampleMin2() {
	fmt.Println(seq.Min2(seq.Skip2(seq.Range2(6), 3)))

	// Output:
	// 3 3 true
}
