package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleMovingAverage() {
	for num := range seq.MovingAverage(seq.Range(0, 30, 5), 3) {
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

func ExampleCenteredMovingAverage() {
	for num := range seq.CenteredMovingAverage(seq.Range(0, 30, 5), 3) {
		if num > 10 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 5
	// 10
}
