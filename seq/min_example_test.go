package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleMin() {
	fmt.Println(seq.Min(seq.Skip(seq.Range(6), 3)))

	// Output:
	// 3 true
}

func ExampleMin2() {
	fmt.Println(seq.Min2(seq.Skip2(seq.Range2(6), 3)))

	// Output:
	// 3 3 true
}

func ExampleMinFunc() {
	fmt.Println(seq.MinFunc(
		seq.Range(6),
		func(item1, item2 int) int { return item1 - item2 },
	))

	// Output:
	// 0 true
}

func ExampleMinFunc2() {
	fmt.Println(seq.MinFunc2(
		seq.Range2(6),
		func(t1, t2 seq.Tuple[int, int]) int { return t1.K - t2.K },
	))

	// Output:
	// 0 0 true
}
