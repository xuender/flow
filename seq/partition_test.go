package seq_test

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/xuender/flow/seq"
)

func TestPartition_1(t *testing.T) {
	t.Parallel()

	var (
		wait  sync.WaitGroup
		count int32
	)

	even, odd := seq.Partition(seq.Range(6), func(num int) bool { return num%2 == 0 })

	wait.Add(2)

	go func() {
		for num := range even {
			if num > 2 {
				break
			}

			atomic.AddInt32(&count, 1)
		}

		wait.Done()
	}()

	go func() {
		for range odd {
			atomic.AddInt32(&count, 1)
		}

		wait.Done()
	}()

	wait.Wait()

	if count != 4 {
		t.Errorf("Partition_1 count %v!=4", count)
	}
}

func TestPartition_2(t *testing.T) {
	t.Parallel()

	var (
		wait  sync.WaitGroup
		count int32
	)

	even, odd := seq.Partition(seq.Range(6), func(num int) bool { return num%2 == 0 })

	wait.Add(2)

	go func() {
		for range even {
			atomic.AddInt32(&count, 1)
		}

		wait.Done()
	}()

	go func() {
		for num := range odd {
			if num > 2 {
				break
			}

			atomic.AddInt32(&count, 1)
		}

		wait.Done()
	}()

	wait.Wait()

	if count != 3 {
		t.Errorf("Partition_2 count %v!=2", count)
	}
}

func TestPartition2_1(t *testing.T) {
	t.Parallel()

	var (
		wait  sync.WaitGroup
		count int32
	)

	even, odd := seq.Partition2(seq.Range2(6), func(key, _ int) bool { return key%2 == 0 })

	wait.Add(2)

	go func() {
		for key := range even {
			if key > 2 {
				break
			}

			atomic.AddInt32(&count, 1)
		}

		wait.Done()
	}()

	go func() {
		for range odd {
			atomic.AddInt32(&count, 1)
		}

		wait.Done()
	}()

	wait.Wait()

	if count != 4 {
		t.Errorf("Partition_1 count %v!=4", count)
	}
}

func TestPartition2_2(t *testing.T) {
	t.Parallel()

	var (
		wait  sync.WaitGroup
		count int32
	)

	even, odd := seq.Partition2(seq.Range2(6), func(key, _ int) bool { return key%2 == 0 })

	wait.Add(2)

	go func() {
		for range even {
			atomic.AddInt32(&count, 1)
		}

		wait.Done()
	}()

	go func() {
		for key := range odd {
			if key > 2 {
				break
			}

			atomic.AddInt32(&count, 1)
		}

		wait.Done()
	}()

	wait.Wait()

	if count != 3 {
		t.Errorf("Partition2_2 count %v!=2", count)
	}
}
