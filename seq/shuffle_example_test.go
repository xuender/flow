package seq_test

import (
	"github.com/xuender/flow/seq"
)

func ExampleShuffle() {
	num := 0
	for range seq.Shuffle(seq.Range(10)) {
		num++
		if num > 3 {
			break
		}
	}

	// Output:
}

func ExampleShuffle2() {
	num := 0
	for range seq.Shuffle2(seq.Range2(10)) {
		num++
		if num > 3 {
			break
		}
	}

	// Output:
}
