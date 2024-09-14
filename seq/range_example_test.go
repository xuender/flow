package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleRange_zero() {
	for num := range seq.Range() {
		fmt.Println(num)
	}

	// Output:
}

func ExampleRange_one() {
	for num := range seq.Range(3) {
		fmt.Println(num)
	}

	// Output:
	// 0
	// 1
	// 2
}

func ExampleRange_oneNegative() {
	for num := range seq.Range(-3) {
		fmt.Println(num)
	}

	// Output:
	// 0
	// -1
	// -2
}

func ExampleRange_two() {
	for num := range seq.Range(3, 5) {
		fmt.Println(num)
	}

	// Output:
	// 3
	// 4
}

func ExampleRange_twoDecrease() {
	for num := range seq.Range(5, 3) {
		fmt.Println(num)
	}

	// Output:
	// 5
	// 4
}

func ExampleRange_twoEqual() {
	for num := range seq.Range(3, 3) {
		fmt.Println(num)
	}

	// Output:
}

func ExampleRange_twoNegative() {
	for num := range seq.Range(-3, -5) {
		fmt.Println(num)
	}

	// Output:
	// -3
	// -4
}

func ExampleRange_twoNegativeIncremental() {
	for num := range seq.Range(-5, -3) {
		fmt.Println(num)
	}

	// Output:
	// -5
	// -4
}

func ExampleRange_three() {
	for num := range seq.Range(3, 10, 3) {
		fmt.Println(num)
	}

	// Output:
	// 3
	// 6
	// 9
}

func ExampleRange_stepDecrease() {
	for num := range seq.Range(-1, -3, -1) {
		fmt.Println(num)
	}

	// Output:
	// -1
	// -2
}

func ExampleRange_stepWrong() {
	for num := range seq.Range(-1, -3, 1) {
		fmt.Println(num)
	}

	// Output:
}

func ExampleRange_repeat() {
	for num := range seq.Range(3, 7, 0) {
		fmt.Println(num)
	}

	// Output:
	// 3
	// 3
	// 3
	// 3
}
