package seq_test

import (
	"fmt"
	"sync"

	"github.com/xuender/flow/seq"
)

func ExamplePartition() {
	even, odd := seq.Partition(seq.Range(6), func(num int) bool { return num%2 == 0 })

	var wait sync.WaitGroup

	wait.Add(2)

	go func() {
		for num := range even {
			fmt.Println("even", num)
		}

		wait.Done()
	}()

	go func() {
		for num := range odd {
			fmt.Println("odd", num)
		}

		wait.Done()
	}()

	wait.Wait()

	// Output:
	// even 0
	// odd 1
	// even 2
	// odd 3
	// even 4
	// odd 5
}

func ExamplePartition2() {
	even, odd := seq.Partition2(seq.Range2(6), func(key, _ int) bool { return key%2 == 0 })

	var wait sync.WaitGroup

	wait.Add(2)

	go func() {
		for key, val := range even {
			fmt.Println("even", key, val)
		}

		wait.Done()
	}()

	go func() {
		for key, val := range odd {
			fmt.Println("odd", key, val)
		}

		wait.Done()
	}()

	wait.Wait()

	// Output:
	// even 0 0
	// odd 1 1
	// even 2 2
	// odd 3 3
	// even 4 4
	// odd 5 5
}
