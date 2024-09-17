package seq_test

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/xuender/flow/seq"
)

func TestDistribute(t *testing.T) {
	t.Parallel()

	seqs := seq.Distribute(seq.Range(100), 5)
	wait := sync.WaitGroup{}

	var count int32

	wait.Add(5)

	for _, input := range seqs {
		go func() {
			for num := range input {
				if num > 10 {
					break
				}

				atomic.AddInt32(&count, 1)
			}

			wait.Done()
		}()
	}

	wait.Wait()

	if count != 11 {
		t.Errorf("Distribute count:%d", count)
	}
}

func TestDistribute2(t *testing.T) {
	t.Parallel()

	seqs := seq.Distribute2(seq.Range2(100), 5)
	wait := sync.WaitGroup{}

	var count int32

	wait.Add(5)

	for _, input := range seqs {
		go func() {
			for num := range input {
				if num > 10 {
					break
				}

				atomic.AddInt32(&count, 1)
			}

			wait.Done()
		}()
	}

	wait.Wait()

	if count != 11 {
		t.Errorf("Distribute2 count:%d", count)
	}
}
