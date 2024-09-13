package seq_test

import (
	"fmt"
	"slices"

	"gitee.com/xuender/flow/seq"
)

func ExampleSum() {
	fmt.Println(seq.Sum(seq.Range(101)))
	fmt.Println(seq.Sum(slices.Values([]string{"a", "b", "c"})))

	// Output:
	// 5050
	// abc
}
