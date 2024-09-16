package main

import (
	"fmt"
	"time"

	"github.com/xuender/flow"
	"github.com/xuender/flow/seq"
)

func main() {
	for num := range flow.Parallel(3,
		seq.Range(100),
		flow.Filter(func(num int) bool { return num%3 == 0 }),
		flow.Skip[int](5),
		flow.Limit[int](4),
		flow.Peek(func(num int) {
			fmt.Println("peek", num)
			time.Sleep(time.Second)
		}),
		flow.Sort[int](),
	) {
		fmt.Println(num)
	}
}
