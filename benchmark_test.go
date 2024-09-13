// nolint: revive
package flow_test

import (
	"slices"
	"testing"

	"github.com/xuender/flow"
	"github.com/xuender/flow/seq"
)

func BenchmarkChain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for range flow.Chain(
			seq.Range(1000),
			flow.Filter(func(num int) bool { return num%3 == 0 }),
			flow.Skip[int](5),
			flow.Limit[int](4),
			flow.Reverse[int](),
		) {
		}
	}
}

func BenchmarkSlices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := []int{}

		for num := range 1000 {
			if num%3 == 0 {
				slice = append(slice, num)
			}
		}

		slice = slice[5 : 5+4]
		slices.Reverse(slice)

		for range slice {
		}
	}
}
