package main

import (
	"fmt"

	"github.com/xuender/flow"
	"github.com/xuender/flow/seq"
)

func main() {
	for num := range flow.Chain(
		seq.Range(3),
		flow.Map(func(num int) int { return num * 2 }),
	) {
		fmt.Println(num)
	}
}
