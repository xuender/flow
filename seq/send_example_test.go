package seq_test

import (
	"fmt"
	"sync"

	"github.com/xuender/flow/seq"
)

func ExampleSend() {
	out := make(chan int)
	group := sync.WaitGroup{}

	group.Add(1)

	go func() {
		group.Wait()
		close(out)
	}()

	go func() {
		seq.Send(seq.Range(3), out)
		group.Done()
	}()

	for num := range out {
		fmt.Println(num)
	}

	// Output:
	// 0
	// 1
	// 2
}

func ExampleSend2() {
	out := make(chan seq.Tuple[int, int])
	group := sync.WaitGroup{}

	group.Add(1)

	go func() {
		group.Wait()
		close(out)
	}()

	go func() {
		seq.Send2(seq.Range2(3), out)
		group.Done()
	}()

	for item := range out {
		fmt.Println(item.K, item.V)
	}

	// Output:
	// 0 0
	// 1 1
	// 2 2
}
