package gootstrap

import (
	"syscall"
)

// Run starts all runners and stops them on SIGTERM or SIGINT.
func Run(runners ...Runner) {
	RunUntilSignal(RunAll(runners...), syscall.SIGTERM, syscall.SIGINT)
}
