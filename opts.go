package l1

import "time"

// Opt
type Opt func(*Runner)

func WithMaxConns(amount int) Opt {
	return func(r *Runner) {
		r.MaxConnections = amount
	}
}

func WithTimeout(s int) Opt {
	return func(r *Runner) {
		r.Timeout = s
	}
}

func WithRunTime(timeInSecs int) Opt {
	return func(r *Runner) {
		r.RunTime = time.Duration(timeInSecs) * time.Second
	}
}
