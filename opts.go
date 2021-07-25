package l1

import "time"

// Opt
type Opt func(*Runner)

func WithMaxConns(amount int) Opt {
	return func(r *Runner) {
		r.MaxConnections = amount
	}
}

func WithTimeout(timeInSecs int) Opt {
	return func(r *Runner) {
		r.Timeout = time.Duration(timeInSecs) * time.Second
	}
}

func WithRunTime(timeInSecs int) Opt {
	return func(r *Runner) {
		r.RunTime = time.Duration(timeInSecs) * time.Second
	}
}

func WithMaxParrellConns(amount int) Opt {
	return func(r *Runner) {
		r.MaxParrellConnections = amount
	}
}

func WithTarget(target string) Opt {
	return func(r *Runner) {
		r.Target = target
	}
}

func WithRunFunc(f F) Opt {
	return func(r *Runner) {
		r.RunFunc = f
	}
}
