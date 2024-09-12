package main

import (
	"log/slog"
	"time"

	"gitee.com/xuender/flow"
)

func main() {
	for num := range flow.Parallel(
		3,
		flow.Range(100),
		flow.Filter(func(num int) bool { return num%3 == 0 }),
		flow.Skip[int](5),
		flow.Limit[int](4),
		flow.Reverse[int](),
		flow.Peek(func(num int) {
			time.Sleep(time.Second)
		}),
	) {
		slog.Info("end", "num", num)
	}
}
