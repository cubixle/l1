package l1

import (
	"sync"
)

type Task struct {
	Target string
	Result *Result
	F      F
}

type pool struct {
	tasks       []*Task
	concurrency int

	tasksChan chan *Task
	wg        sync.WaitGroup
}

func newPool(tasks []*Task, concurrency int) *pool {
	return &pool{
		tasks:       tasks,
		concurrency: concurrency,
		tasksChan:   make(chan *Task),
	}
}

// run will run all work within the pool.
func (p *pool) run() {
	for i := 0; i < p.concurrency; i++ {
		go p.work()
	}

	p.wg.Add(len(p.tasks))
	for _, t := range p.tasks {
		p.tasksChan <- t
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
