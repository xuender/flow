package seq_test

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/xuender/flow/seq"
)

type Num int

func (p Num) String() string {
	return strconv.Itoa(int(p))
}

func ExampleJoin() {
	fmt.Println(seq.Join(seq.Range(3), ","))
	fmt.Println(seq.Join(slices.Values([]string{"a", "b", "c"}), ","))
	fmt.Println(seq.Join(seq.Map(seq.Range(3), func(num int) Num { return Num(num) }), ","))

	// Output:
	// 0,1,2
	// a,b,c
	// 0,1,2
}
