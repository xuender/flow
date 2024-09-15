package seq_test

import (
	"fmt"

	"github.com/xuender/flow/seq"
)

func ExampleChunk() {
	seqs := seq.Chunk(seq.Range(5), 2)

	for _, items := range seqs {
		for num := range items {
			fmt.Println(num)
		}
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func ExampleChunk2() {
	seqs := seq.Chunk2(seq.Range2(5), 2)

	for _, items := range seqs {
		for key, val := range items {
			if key > 2 {
				break
			}

			fmt.Println(key, val)
		}
	}

	// Output:
	// 0 0
	// 1 1
	// 2 2
}
