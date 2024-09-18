package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleWindow() {
	for items := range seq.Window(3, seq.Range(5)) {
		if items[0] > 1 {
			break
		}

		fmt.Println(items)
	}

	// Output:
	// [0]
	// [0 1]
	// [0 1 2]
	// [1 2 3]
}

func ExampleWindow2() {
	for items := range seq.Window2(3, seq.Range2(5)) {
		if items[0].K > 1 {
			break
		}

		fmt.Println(items)
	}

	// Output:
	// [{0 0}]
	// [{0 0} {1 1}]
	// [{0 0} {1 1} {2 2}]
	// [{1 1} {2 2} {3 3}]
}
