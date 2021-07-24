package l1

import "time"

// Runner
type Runner struct {
	MaxConnections int
	Timeout        int
	RunTime        time.Duration
}

func NewRunner(opts ...Opt) {
	r := &Runner{
		RunTime: 60 * time.Second,
	}
	for _, o := range opts {
		o(r)
	}
}

// F defines the function type for runners.
type F func(url string) error
