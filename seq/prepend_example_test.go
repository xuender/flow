package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExamplePrepend() {
	for num := range seq.Prepend(seq.Range(1, 3), 4, 5) {
		if num == 2 {
			break
		}

		fmt.Println(num)
	}

	for num := range seq.Prepend(seq.Range(1, 3), 7, 8) {
		if num < 8 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 4
	// 5
	// 1
}

func ExamplePrepend2() {
	for key, val := range seq.Prepend2(seq.Range2(1, 3), seq.T(4, 8), seq.T(5, 10)) {
		if key < 3 {
			break
		}

		fmt.Println(key, val)
	}

	for key, val := range seq.Prepend2(seq.Range2(1, 3), seq.T(4, 8), seq.T(5, 10)) {
		if key > 4 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 4 8
	// 5 10
	// 4 8
}
