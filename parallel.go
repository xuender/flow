package flow

import (
	"iter"
	"sync"

	"gitee.com/xuender/flow/seq"
)

type parallel[E any] struct {
	wait   sync.WaitGroup
	size   int
	output chan E
	chans  []chan E
}

func Parallel[E any](size int, items iter.Seq[E], steps ...Step[E, E]) iter.Seq[E] {
	par := parallel[E]{
		size:   size,
		output: make(chan E),
		chans:  make([]chan E, size),
	}

	for idx := range size {
		par.chans[idx] = make(chan E)

		par.wait.Add(1)

		go par.call(idx, steps)
	}

	go par.consumer(items)

	go par.close()

	return seq.Chan(par.output)
}

func (p *parallel[E]) close() {
	p.wait.Wait()
	close(p.output)
}

func (p *parallel[E]) consumer(seq iter.Seq[E]) {
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

func (p *parallel[E]) send(idx int, item E) bool {
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

func (p *parallel[E]) call(idx int, steps []Step[E, E]) {
	items := seq.Chan(p.chans[idx])

	for _, step := range steps {
		items = step(items)
	}

	for item := range items {
		p.output <- item
	}

	p.wait.Done()
}
