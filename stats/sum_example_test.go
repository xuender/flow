package stats_test

import (
	"fmt"

	"github.com/xuender/flow/stats"
)

func ExampleSum() {
	fmt.Println(stats.Sum([]int{1, 2, 3, 4, 5}))
	fmt.Println(stats.Sum([]string{"a", "b", "c"}))

	// Output:
	// 15
	// abc
}
