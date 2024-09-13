package main

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/xuender/flow"
	"github.com/xuender/flow/seq"
)

func main() {
	items := flow.Parallel(3,
		seq.Range(100),
		flow.Filter(func(num int) bool { return num%3 == 0 }),
		flow.Peek(func(num int) {
			fmt.Println("peek", num)
			time.Sleep(time.Second)
		}))

	for num := range flow.Chain(
		items,
		flow.Skip[int](5),
		flow.Limit[int](4),
		flow.Reverse[int](),
	) {
		slog.Info("end", "num", num)
	}
}
