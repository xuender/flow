package flow

import (
	"iter"
	"sync"

	"gitee.com/xuender/flow/seq"
)

type controller[E any] struct {
	wait   sync.WaitGroup
	size   int
	output chan E
	chans  []chan E
}

func Parallel[E any](size int, items iter.Seq[E], steps ...Step[E]) iter.Seq[E] {
	ctrl := controller[E]{
		size:   size,
		output: make(chan E),
		chans:  make([]chan E, size),
	}

	for idx := range size {
		ctrl.chans[idx] = make(chan E)

		ctrl.wait.Add(1)

		go ctrl.run(idx, steps)
	}

	go ctrl.consum(items)

	go ctrl.close()

	return seq.Chan(ctrl.output)
}

func (p *controller[E]) close() {
	p.wait.Wait()
	close(p.output)
}

func (p *controller[E]) consum(seq iter.Seq[E]) {
	idx := 0
	for item := range seq {
		if p.send(idx%p.size, item) {
			break
		}

		idx++
	}

	for idx := range p.size {
		if p.chans[idx] != nil {
			close(p.chans[idx])
		}
	}
}

func (p *controller[E]) send(idx int, item E) bool {
	if p.chans[idx] == nil {
		return true
	}

	defer func() {
		if err := recover(); err != nil {
			p.chans[idx] = nil
		}
	}()

	p.chans[idx] <- item

	return false
}

func (p *controller[E]) run(idx int, steps []Step[E]) {
	items := seq.Chan(p.chans[idx])

	for _, step := range steps {
		items = step(items)
	}

	for item := range items {
		p.output <- item
	}

	p.wait.Done()
}
