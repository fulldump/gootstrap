package gootstrap

import "testing"

func TestRunUntilSignalWithNilRunnerDoesNotPanic(t *testing.T) {
	RunUntilSignal(nil)
}

func TestRunUntilSignalWithNilFunctionsDoesNotPanic(t *testing.T) {
	r := func() (func() error, func() error) {
		return nil, nil
	}

	RunUntilSignal(r)
}
