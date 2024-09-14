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

func ExampleMap2() {
	for key, val := range seq.Map2(
		seq.Range2(6),
		func(key, val int) (int, float64) { return key, float64(val) * 1.35 },
	) {
		if key > 2 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 0 0
	// 1 1.35
	// 2 2.7
}
