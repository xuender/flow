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

func ExampleRange2_zero() {
	for key, val := range seq.Range2() {
		fmt.Println(key, val)
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

func ExampleRange2_one() {
	for key, val := range seq.Range2(3) {
		fmt.Println(key, val)
	}

	// Output:
	// 0 0
	// 1 1
	// 2 2
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

func ExampleRange2_oneNegative() {
	for key, val := range seq.Range2(-3) {
		fmt.Println(key, val)
	}

	// Output:
	// 0 0
	// 1 -1
	// 2 -2
}

func ExampleRange_two() {
	for num := range seq.Range(3, 5) {
		fmt.Println(num)
	}

	// Output:
	// 3
	// 4
}

func ExampleRange2_two() {
	for key, val := range seq.Range2(3, 5) {
		fmt.Println(key, val)
	}

	// Output:
	// 0 3
	// 1 4
}

func ExampleRange_twoDecrease() {
	for num := range seq.Range(5, 3) {
		fmt.Println(num)
	}

	// Output:
	// 5
	// 4
}

func ExampleRange2_twoDecrease() {
	for key, val := range seq.Range2(5, 3) {
		fmt.Println(key, val)
	}

	// Output:
	// 0 5
	// 1 4
}

func ExampleRange_twoEqual() {
	for num := range seq.Range(3, 3) {
		fmt.Println(num)
	}

	// Output:
}

func ExampleRange2_twoEqual() {
	for key, val := range seq.Range2(3, 3) {
		fmt.Println(key, val)
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

func ExampleRange2_twoNegative() {
	for key, val := range seq.Range2(-3, -5) {
		fmt.Println(key, val)
	}

	// Output:
	// 0 -3
	// 1 -4
}

func ExampleRange_twoNegativeIncremental() {
	for num := range seq.Range(-5, -3) {
		fmt.Println(num)
	}

	// Output:
	// -5
	// -4
}

func ExampleRange2_twoNegativeIncremental() {
	for key, val := range seq.Range2(-5, -3) {
		fmt.Println(key, val)
	}

	// Output:
	// 0 -5
	// 1 -4
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

func ExampleRange2_three() {
	for key, val := range seq.Range2(3, 10, 3) {
		fmt.Println(key, val)
	}

	// Output:
	// 0 3
	// 1 6
	// 2 9
}

func ExampleRange_stepDecrease() {
	for num := range seq.Range(-1, -3, -1) {
		fmt.Println(num)
	}

	// Output:
	// -1
	// -2
}

func ExampleRange2_stepDecrease() {
	for key, val := range seq.Range2(-1, -3, -1) {
		fmt.Println(key, val)
	}

	// Output:
	// 0 -1
	// 1 -2
}

func ExampleRange_stepWrong() {
	for num := range seq.Range(-1, -3, 1) {
		fmt.Println(num)
	}

	// Output:
}

func ExampleRange2_stepWrong() {
	for key, val := range seq.Range2(-1, -3, 1) {
		fmt.Println(key, val)
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

func ExampleRange2_repeat() {
	for key, val := range seq.Range2(3, 7, 0) {
		fmt.Println(key, val)
	}

	// Output:
	// 0 3
	// 1 3
	// 2 3
	// 3 3
}
