package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleT() {
	fmt.Println(seq.T(1, "a"))

	// Output:
	// {1 a}
}

func ExampleTuples() {
	fmt.Println(seq.Tuples(seq.Range2(2)))

	// Output:
	// [{0 0} {1 1}]
}
