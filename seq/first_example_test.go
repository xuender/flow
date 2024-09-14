package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleFirst() {
	fmt.Println(seq.First(seq.Range(100)))
	fmt.Println(seq.First(seq.Range(0)))

	// Output:
	// 0 true
	// 0 false
}
