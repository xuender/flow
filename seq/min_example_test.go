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
