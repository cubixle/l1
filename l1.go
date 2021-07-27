package l1

import (
	"time"
)

// F defines the function type for runners.
type F func(target string) *Result

// Runner
type Runner struct {
	MaxConnections        int
	MaxParrellConnections int
	Timeout               time.Duration
	RunTime               time.Duration
	RunFunc               F
	Target                string
	results               []*Result
}

func NewRunner(opts ...Opt) (*Runner, error) {
	r := &Runner{
		RunTime:               60 * time.Second,
		Timeout:               30 * time.Second,
		MaxParrellConnections: 10,
		MaxConnections:        10,
	}

	for _, o := range opts {
		o(r)
	}

	if r.Target == "" {
		return nil, ErrNoTarget
	}

	return r, nil
}

func (r *Runner) SetOpt(o Opt) {
	o(r)
}

func (r *Runner) Execute() {
	tasks := []*Task{}
	for i := 0; i < r.MaxConnections; i++ {
		tasks = append(tasks, &Task{Target: r.Target, F: r.RunFunc})
	}
	// create the pool and process the tasks.
	pool := newPool(tasks, r.MaxParrellConnections)
	// the tasks are updated in memory so we don't expect a return here.
	pool.run()
	for _, t := range tasks {
		r.results = append(r.results, t.Result)
	}
}

func (r *Runner) Results() *results {
	res := &results{
		Results: r.results,
		Target:  r.Target,
	}
	return res
}
