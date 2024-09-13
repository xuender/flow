package flow_test

import (
	"slices"
	"testing"

	"gitee.com/xuender/flow"
	"gitee.com/xuender/flow/seq"
)

func BenchmarkChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seq.Emit(
			flow.Chain(
				seq.Range(100),
				flow.Filter(func(num int) bool { return num%3 == 0 }),
				flow.Skip[int](5),
				flow.Limit[int](4),
				flow.Reverse[int](),
			))
	}
}

func BenchmarkParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seq.Emit(
			flow.Parallel(3,
				seq.Range(100),
				flow.Filter(func(num int) bool { return num%3 == 0 }),
				flow.Skip[int](5),
				flow.Limit[int](4),
				flow.Reverse[int](),
			))
	}
}

func BenchmarkSlices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := []int{}

		for num := range 100 {
			if num%3 == 0 {
				slice = append(slice, num)
			}
		}

		slice = slice[5 : 5+4]
		slices.Reverse(slice)
	}
}
