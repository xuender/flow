package seq_test

import (
	"testing"

	"github.com/xuender/flow/seq"
)

func TestToChans(t *testing.T) {
	t.Parallel()

	chans := seq.ToChans(seq.Range(10), 3)
	sum := 0

	for _, cha := range chans {
		for item := range cha {
			if item > 2 {
				break
			}

			sum += item
		}
	}

	if sum != 3 {
		t.Errorf("ToChans count:%d", sum)
	}
}

func TestToChans2(t *testing.T) {
	t.Parallel()

	chans := seq.ToChans2(seq.Range2(10), 3)
	sum := 0

	for _, cha := range chans {
		for item := range cha {
			if item.K > 2 {
				break
			}

			sum += item.K
		}
	}

	if sum != 3 {
		t.Errorf("ToChans2 count:%d", sum)
	}
}
