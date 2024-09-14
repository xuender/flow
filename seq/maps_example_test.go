package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleKeys() {
	for key := range seq.Keys(seq.Range2(3)) {
		if key > 1 {
			break
		}

		fmt.Println(key)
	}

	// Output:
	// 0
	// 1
}

func ExampleValues() {
	for val := range seq.Values(seq.Range2(3)) {
		if val > 1 {
			break
		}

		fmt.Println(val)
	}

	// Output:
	// 0
	// 1
}
