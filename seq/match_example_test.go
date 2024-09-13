// nolint: dupword
package seq_test

import (
	"fmt"

	"gitee.com/xuender/flow/seq"
)

func ExampleAnyMatch() {
	fmt.Println(seq.AnyMatch(seq.Range(3), func(num int) bool { return num > 1 }))
	fmt.Println(seq.AnyMatch(seq.Range(3), func(num int) bool { return num > 3 }))
	fmt.Println(seq.AnyMatch(seq.Range(0), func(num int) bool { return num > 3 }))

	// Output:
	// true
	// false
	// false
}

func ExampleAllMatch() {
	fmt.Println(seq.AllMatch(seq.Range(3), func(num int) bool { return num > 1 }))
	fmt.Println(seq.AllMatch(seq.Range(3), func(num int) bool { return num < 3 }))
	fmt.Println(seq.AllMatch(seq.Range(0), func(num int) bool { return num < 3 }))

	// Output:
	// false
	// true
	// false
}

func ExampleNoneMatch() {
	fmt.Println(seq.NoneMatch(seq.Range(3), func(num int) bool { return num > 1 }))
	fmt.Println(seq.NoneMatch(seq.Range(3), func(num int) bool { return num > 3 }))
	fmt.Println(seq.NoneMatch(seq.Range(0), func(num int) bool { return num > 3 }))

	// Output:
	// false
	// true
	// true
}
