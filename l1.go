package l1

import (
	"time"
)

// F defines the function type for runners.
type F func(target string) error

// Runner
type Runner struct {
	MaxConnections        int
	MaxParrellConnections int
	Timeout               time.Duration
	RunTime               time.Duration
	RunFunc               F
	Target                string
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
