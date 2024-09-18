package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleMovingAvg() {
	for num := range seq.MovingAvg(seq.Range(0, 30, 5), 3) {
		if num > 12 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 0
	// 2
	// 5
	// 10
}

func ExampleCenteredMovingAvg() {
	for num := range seq.CenteredMovingAvg(seq.Range(0, 30, 5), 3) {
		if num > 10 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 5
	// 10
}
