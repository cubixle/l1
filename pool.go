package l1

import (
	"sync"
	"time"
)

type Task struct {
	Target string
	Result *Result
	F      F
}

type pool struct {
	tasks          []*Task
	concurrency    int
	rampUpDuration int

	tasksChan chan *Task
	wg        sync.WaitGroup
}

func newPool(tasks []*Task, concurrency, rampUpDuration int) *pool {
	return &pool{
		tasks:          tasks,
		concurrency:    concurrency,
		tasksChan:      make(chan *Task),
		rampUpDuration: rampUpDuration,
	}
}

// run will run all work within the pool.
func (p *pool) run() {
	for i := 0; i < p.concurrency; i++ {
		go p.work()
	}

	for _, t := range p.tasks {
		p.tasksChan <- t
		p.wg.Add(1)

		time.Sleep(time.Duration((1000 * p.rampUpDuration) * int(time.Millisecond)))
	}

	close(p.tasksChan)

	p.wg.Wait()
}

func (p *pool) work() {
	for t := range p.tasksChan {
		if t.F == nil {
			continue
		}
		res := t.F(t.Target)
		t.Result = res
		p.wg.Done()
	}
}
