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
	RampUp                int
	results               []*Result
}

// NewRunner creates a new Runner...
// There a bunch of defaults that will be set that can be overwritten
// with the Options (Opt) found in opts.go.
//
// The only error that is returned currently is if no target has been set
// with the WithTarget option func.
func NewRunner(opts ...Opt) (*Runner, error) {
	r := &Runner{
		MaxParrellConnections: 10,
		MaxConnections:        10,
		RampUp:                0,
	}

	for _, o := range opts {
		o(r)
	}

	if r.Target == "" {
		return nil, ErrNoTarget
	}

	return r, nil
}

// SetOpt gives the ability to change the configuration of Runner in a
// standardised way after NewRunner has been called.
func (r *Runner) SetOpt(o Opt) {
	o(r)
}

// Execute runs a worker pool and then fires off requests until MaxConnections
// has been reached.
//
// Execute will block until all requests have been completed.
//
// All results are aggregated and are accessible via the .Results() method
// once the requests have completed.
func (r *Runner) Execute() {
	tasks := []*Task{}
	for i := 0; i < r.MaxConnections; i++ {
		tasks = append(tasks, &Task{Target: r.Target, F: r.RunFunc})
	}

	waitDuration := 0
	if r.RampUp > 0 {
		waitDuration = r.RampUp / len(tasks)
	}

	// create the pool and process the tasks.
	pool := newPool(tasks, r.MaxParrellConnections, waitDuration)

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
