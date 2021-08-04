package l1

import "fmt"

var (
	// ErrNoTarget is the error used when no target has been set
	// in the Runner.
	ErrNoTarget = fmt.Errorf("no target has been set")
)
