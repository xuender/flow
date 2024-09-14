package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleAppend() {
	for num := range seq.Append(seq.Range(1, 3), 4, 5) {
		if num > 4 {
			break
		}

		fmt.Println(num)
	}

	for num := range seq.Append(seq.Range(1, 3), 4, 5) {
		if num > 1 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 1
	// 2
	// 4
	// 1
}

func ExampleAppend2() {
	for key, val := range seq.Append2(seq.Range2(1, 3), seq.T(4, 8), seq.T(5, 10)) {
		if key > 4 {
			break
		}

		fmt.Println(key, val)
	}

	for key, val := range seq.Append2(seq.Range2(1, 3), seq.T(4, 8), seq.T(5, 10)) {
		if key > 0 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 0 1
	// 1 2
	// 4 8
	// 0 1
}
