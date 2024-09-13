package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleJoin() {
	fmt.Println(seq.Join(seq.Range(3), ","))

	// Output:
	// 0,1,2
}
