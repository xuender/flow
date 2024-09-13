package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleFilter() {
	fmt.Println(seq.Count(seq.Filter(seq.Range(6), func(i int) bool { return i%2 == 0 })))

	// Output:
	// 3
}
