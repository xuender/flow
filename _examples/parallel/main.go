package main

import (
	"fmt"
	"time"

	"github.com/xuender/flow"
	"github.com/xuender/flow/seq"
)

func main() {
	start := time.Now()

	for num := range flow.Chain(
		seq.Range(100),
		flow.Filter(func(num int) bool { return num%3 == 0 }),
		flow.Skip[int](5),
		flow.Limit[int](4),
		flow.Peek(func(num int) {
			fmt.Println("peek", num)
			time.Sleep(time.Second)
		}),
		flow.Reverse[int](),
	) {
		fmt.Println(num)
	}

	fmt.Printf("Chain: %v\n", time.Since(start))
	start = time.Now()

	for num := range flow.Parallel(3,
		seq.Range(100),
		flow.Filter(func(num int) bool { return num%3 == 0 }),
		flow.Skip[int](5),
		flow.Limit[int](4),
		flow.Peek(func(num int) {
			fmt.Println("peek", num)
			time.Sleep(time.Second)
		}),
		flow.Reverse[int](),
	) {
		fmt.Println(num)
	}

	fmt.Printf("Parallel: %v\n", time.Since(start))
}
