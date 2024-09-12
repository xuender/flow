package main

import (
	"fmt"

	"gitee.com/xuender/flow"
)

func main() {
	flow.Each(
		flow.Chain(
			flow.Range(100),
			flow.Filter(func(num int) bool { return num%3 == 0 }),
			flow.Skip[int](5),
			flow.Limit[int](4),
			flow.Reverse[int](),
		),
		func(num int) bool {
			fmt.Println(num)

			return true
		},
	)
}
