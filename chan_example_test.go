package flow_test

import (
	"fmt"
	"sync"

	"gitee.com/xuender/flow"
	"gitee.com/xuender/flow/seq"
)

func ExampleChans() {
	sum := 0
	chans := flow.Chans(seq.Range(10), 3)
	wait := sync.WaitGroup{}
	wait.Add(3)

	for idx, cha := range chans {
		go func(input chan int, num int) {
			for item := range input {
				sum += item
				sum += num
			}

			wait.Done()
		}(cha, idx)
	}

	wait.Wait()
	fmt.Println(sum)

	// Output:
	// 56
}
