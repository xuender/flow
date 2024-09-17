package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleTakeWhile() {
	for num := range seq.TakeWhile(seq.Range(10), func(num int) bool {
		return num < 3
	}) {
		fmt.Println(num)
	}

	for num := range seq.TakeWhile(seq.Range(10), func(num int) bool {
		return num < 3
	}) {
		if num > 1 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 0
	// 1
	// 2
	// 0
	// 1
}

func ExampleDropWhile() {
	for num := range seq.DropWhile(seq.Range(10), func(num int) bool {
		return num < 5
	}) {
		if num > 8 {
			break
		}

		fmt.Println(num)
	}

	// Output:
	// 5
	// 6
	// 7
	// 8
}

func ExampleTakeWhile2() {
	for key, val := range seq.TakeWhile2(seq.Range2(10), func(key, _ int) bool {
		return key < 3
	}) {
		fmt.Println(key, val)
	}

	for key, val := range seq.TakeWhile2(seq.Range2(10), func(key, _ int) bool {
		return key < 3
	}) {
		if key > 1 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 0 0
	// 1 1
	// 2 2
	// 0 0
	// 1 1
}

func ExampleDropWhile2() {
	for key, val := range seq.DropWhile2(seq.Range2(10), func(key, _ int) bool {
		return key < 5
	}) {
		if key > 8 {
			break
		}

		fmt.Println(key, val)
	}

	// Output:
	// 5 5
	// 6 6
	// 7 7
	// 8 8
}
