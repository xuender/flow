package seq_test

import (
	"maps"
	"testing"

	"github.com/xuender/flow/seq"
)

const _size = 5000

func genMap() map[int]int {
	ret := make(map[int]int)
	for i := range _size {
		ret[i] = i
	}

	return ret
}

func pass(_ int) {}

func BenchmarkIterSlice(b *testing.B) {
	val := genMap()

	b.ResetTimer()

	for range b.N {
		_ = seq.Collect(maps.Keys(val), _size)
	}
}

func BenchmarkIterFor(b *testing.B) {
	val := genMap()

	b.ResetTimer()

	for range b.N {
		for key := range maps.Keys(val) {
			pass(key)
		}
	}
}

func BenchmarkGeneric(b *testing.B) {
	val := genMap()

	b.ResetTimer()

	for range b.N {
		_ = genericKeys(val)
	}
}

func BenchmarkGenericFor(b *testing.B) {
	val := genMap()

	b.ResetTimer()

	for range b.N {
		for key := range genericKeys(val) {
			pass(key)
		}
	}
}

func genericKeys[K comparable, V any](in map[K]V) []K {
	result := make([]K, 0, len(in))

	for k := range in {
		result = append(result, k)
	}

	return result
}

func BenchmarkRaw(b *testing.B) {
	val := genMap()

	b.ResetTimer()

	for range b.N {
		_ = rawKeys(val)
	}
}

func BenchmarkRawFor(b *testing.B) {
	val := genMap()

	b.ResetTimer()

	for range b.N {
		for key := range rawKeys(val) {
			pass(key)
		}
	}
}

func rawKeys(maps map[int]int) []int {
	keys := make([]int, 0, len(maps))

	for key := range maps {
		keys = append(keys, key)
	}

	return keys
}
