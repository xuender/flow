package flow_test

import (
	"testing"

	"github.com/xuender/flow"
	"github.com/xuender/flow/seq"
)

func TestParallel(t *testing.T) {
	t.Parallel()

	count := seq.Count(flow.Parallel(2,
		seq.Range(100),
		flow.Limit[int](8),
		flow.Filter(func(num int) bool { return num%3 == 0 }),
	))

	if count != 3 {
		t.Errorf("Parallel count %d != 3", count)
	}
}

func TestParallel_break(t *testing.T) {
	t.Parallel()

	count := 0

	for num := range flow.Parallel(2,
		seq.Range(100),
		flow.Limit[int](8),
		flow.Filter(func(num int) bool { return num%3 == 0 }),
	) {
		if num > 3 {
			break
		}

		count++
	}

	if count != 2 {
		t.Errorf("Parallel break count %d != 2", count)
	}
}

func TestParallel2(t *testing.T) {
	t.Parallel()

	count := seq.Count2(flow.Parallel2(2,
		seq.Range2(100),
		flow.Limit2[int, int](8),
		flow.Filter2(func(key, _ int) bool { return key%3 == 0 }),
	))

	if count != 3 {
		t.Errorf("Parallel2 count %d != 3", count)
	}
}

func TestParallel2_break(t *testing.T) {
	t.Parallel()

	count := 0

	for num := range flow.Parallel2(2,
		seq.Range2(100),
		flow.Limit2[int, int](8),
		flow.Filter2(func(key, _ int) bool { return key%3 == 0 }),
	) {
		if num > 3 {
			break
		}

		count++
	}

	if count != 2 {
		t.Errorf("Parallel2 break count %d != 2", count)
	}
}
