package l1

import (
	"fmt"
)

// Start starts the runner.
func (r *Runner) Start() error {
	jobChan := make(chan string, r.MaxParrellConnections)
	// resultsChan := make(chan struct{})

	for i := 0; i < r.MaxParrellConnections; i++ {
		go func(jobChan chan string) {
			for t := range jobChan {
				r.RunFunc(t)
			}
		}(jobChan)
	}

	return fmt.Errorf("unimplemented")
}
