package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleMap() {
	int2float64 := seq.Map(
		seq.Range(6),
		func(num int) float64 {
			return float64(num) * 1.35
		},
	)

	fmt.Println(seq.Max(int2float64))

	// Output:
	// 6.75 true
}
