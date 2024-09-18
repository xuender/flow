package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleClone() {
	for _, items := range seq.Clone(seq.Range(2), 3) {
		for num := range items {
			if num > 0 {
				break
			}

			fmt.Println(num)
		}
	}

	// Output:
	// 0
	// 0
	// 0
}

func ExampleClone2() {
	for _, items := range seq.Clone2(seq.Range2(2), 3) {
		for key, val := range items {
			if key > 0 {
				break
			}

			fmt.Println(key, val)
		}
	}

	// Output:
	// 0 0
	// 0 0
	// 0 0
}
