package seq_test

import (
	"iter"
	"reflect"
	"slices"
	"testing"

	"github.com/xuender/flow/seq"
)

func TestCenteredMovingAverage1(t *testing.T) {
	t.Parallel()

	sum := 0
	for num := range seq.CenteredMovingAverage(seq.Range(0, 30, 5), 3) {
		sum += num
	}

	if sum != 92 {
		t.Errorf("CenteredMovingAverage sum:%d", sum)
	}
}

func TestCenteredMovingAverage_break(t *testing.T) {
	t.Parallel()

	sum := 0

	for num := range seq.CenteredMovingAverage(seq.Range(0, 30, 5), 3) {
		if num > 20 {
			break
		}

		sum += num
	}

	if sum != 70 {
		t.Errorf("CenteredMovingAverage sum:%d", sum)
	}
}

func TestMovingAverage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input iter.Seq[int]
		want  []int
	}{
		{"0", seq.Range(1, 1), nil},
		{"1", seq.Range(1, 2), []int{1}},
		{"2", seq.Range(1, 3), []int{1, 1}},
		{"3", seq.Range(1, 4), []int{1, 1, 2}},
		{"4", seq.Range(1, 5), []int{1, 1, 2, 3}},
		{"5", seq.Range(1, 6), []int{1, 1, 2, 3, 4}},
		{"6", seq.Range(1, 7), []int{1, 1, 2, 3, 4, 5}},
	}
	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			t.Parallel()

			if got := slices.Collect(seq.MovingAverage(item.input, 3)); !reflect.DeepEqual(got, item.want) {
				t.Errorf("MovingAverage() = %v, want %v", got, item.want)
			}
		})
	}
}

func TestCenteredMovingAverage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input iter.Seq[int]
		want  []int
	}{
		{"0", seq.Range(1, 1), nil},
		{"1", seq.Range(1, 2), []int{1}},
		{"2", seq.Range(1, 3), []int{1, 2}},
		{"3", seq.Range(1, 4), []int{2, 2, 2}},
		{"4", seq.Range(1, 5), []int{2, 3, 3, 3}},
		{"5", seq.Range(1, 6), []int{2, 3, 4, 4, 4}},
		{"6", seq.Range(1, 7), []int{2, 3, 4, 5, 5, 5}},
		{"7", seq.Range(1, 8), []int{2, 3, 4, 5, 6, 6, 6}},
	}
	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			t.Parallel()

			if got := slices.Collect(seq.CenteredMovingAverage(item.input, 3)); !reflect.DeepEqual(got, item.want) {
				t.Errorf("CenteredMovingAverage() = %v, want %v", got, item.want)
			}
		})
	}
}
