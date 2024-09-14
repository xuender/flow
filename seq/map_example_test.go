package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleMap() {
	for num := range seq.Map(
		seq.Range(6),
		func(num int) float64 { return float64(num) * 1.35 },
	) {
		if num > 3.0 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 0
	// 1.35
	// 2.7
}
