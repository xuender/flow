package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleSlice() {
	fmt.Println(seq.Slice(seq.Range(3), 3))

	// Output:
	// [0 1 2]
}
