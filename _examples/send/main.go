package main

import (
	"fmt"
	"sync"

	"github.com/xuender/flow/seq"
)

func main() {
	out := make(chan int)
	group := sync.WaitGroup{}
	group.Add(2)

	go func() {
		group.Wait()
		close(out)
	}()

	go func() {
		seq.Send(seq.Range(5), out)
		group.Done()
	}()

	go func() {
		seq.Send(seq.Range(10, 15), out)
		group.Done()
	}()

	for num := range out {
		fmt.Println(num)
	}
}
