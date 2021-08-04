package l1

// Opt
type Opt func(*Runner)

// WithMaxConns will set the maximum amount of connections
// you want the runners to make.
//
// If this is not passed to the NewRunner() the default will be
// set to 10.
func WithMaxConns(amount int) Opt {
	return func(r *Runner) {
		r.MaxConnections = amount
	}
}

// WithMaxParrellConns will set the max parrell connections
// you want the runners to use.
//
// If this is not passed to the NewRunner() the default will be
// set to 10.
func WithMaxParrellConns(amount int) Opt {
	return func(r *Runner) {
		r.MaxParrellConnections = amount
	}
}

// WithTarget sets the target, the target will be passed to the runner
// function.
func WithTarget(target string) Opt {
	return func(r *Runner) {
		r.Target = target
	}
}

// WithRunFunc will set the function used by the requests pool.
func WithRunFunc(f F) Opt {
	return func(r *Runner) {
		r.RunFunc = f
	}
}

// WithRampUp will set the amount of time between each ramp up
// stage.
func WithRampUp(seconds int) Opt {
	return func(r *Runner) {
		r.RampUp = seconds
	}
}
