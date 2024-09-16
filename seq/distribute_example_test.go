package seq_test

import (
	"fmt"
	"sync"

	"github.com/xuender/flow/seq"
)

func ExampleDistribute() {
	seqs := seq.Distribute(seq.Range(100), 5)
	wait := sync.WaitGroup{}
	count := 0

	wait.Add(5)

	for _, input := range seqs {
		go func() {
			for range input {
				count++
			}

			wait.Done()
		}()
	}

	wait.Wait()
	fmt.Println(count)

	// Output:
	// 100
}

func ExampleDistribute_break() {
	seqs := seq.Distribute(seq.Range(100), 5)
	wait := sync.WaitGroup{}
	count := 0

	wait.Add(5)

	for _, input := range seqs {
		go func() {
			for num := range input {
				if num > 10 {
					break
				}

				count++
			}

			wait.Done()
		}()
	}

	wait.Wait()
	fmt.Println(count)

	// Output:
	// 11
}

func ExampleDistribute2() {
	seqs := seq.Distribute2(seq.Range2(100), 5)
	wait := sync.WaitGroup{}
	count := 0

	wait.Add(5)

	for _, input := range seqs {
		go func() {
			for range input {
				count++
			}

			wait.Done()
		}()
	}

	wait.Wait()
	fmt.Println(count)

	// Output:
	// 100
}

func ExampleDistribute2_break() {
	seqs := seq.Distribute2(seq.Range2(100), 5)
	wait := sync.WaitGroup{}
	count := 0

	wait.Add(5)

	for _, input := range seqs {
		go func() {
			for num := range input {
				if num > 10 {
					break
				}

				count++
			}

			wait.Done()
		}()
	}

	wait.Wait()
	fmt.Println(count)

	// Output:
	// 11
}
