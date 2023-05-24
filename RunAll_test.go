package gootstrap

import (
	"log"
	"os"
	"testing"
	"time"
)

// TestRunAll tests if all runners are started and stopped when the RunAll's
// start and stop functions are called.
func TestRunAll(t *testing.T) {
	log.SetOutput(os.Stdout)

	var startOneWasCalled, stopOneWasCalled bool

	runnerOne := func() (func() error, func() error) {
		return func() error {
				startOneWasCalled = true
				return nil
			}, func() error {
				stopOneWasCalled = true
				return nil
			}
	}

	var startTwoWasCalled, stopTwoWasCalled bool

	runnerTwo := func() (func() error, func() error) {
		return func() error {
				startTwoWasCalled = true
				return nil
			}, func() error {
				stopTwoWasCalled = true
				return nil
			}
	}

	allRunner := RunAll(runnerOne, runnerTwo)

	start, stop := allRunner()

	err := start()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !startOneWasCalled {
		t.Error("Expected runnerOne start to be called")
	}

	if !startTwoWasCalled {
		t.Error("Expected runnerTwo start to be called")
	}

	err = stop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !stopOneWasCalled {
		t.Error("Expected runnerOne stop to be called")
	}

	if !stopTwoWasCalled {
		t.Error("Expected runnerTwo stop to be called")
	}
}

// TestRunAllWithWait tests if the RunAll's start and stop functions wait for
// all runners to finish starting and stopping.
func TestRunAllWithWait(t *testing.T) {
	log.SetOutput(os.Stdout)

	runner := func() (func() error, func() error) {
		return func() error {
				time.Sleep(1 * time.Second)
				return nil
			}, func() error {
				time.Sleep(1 * time.Second)
				return nil
			}
	}

	allRunner := RunAll(runner, runner)

	start, stop := allRunner()

	startTime := time.Now()
	start()
	stop()
	duration := time.Since(startTime)

	if duration < 2*time.Second {
		t.Error("Expected RunAll to wait for all runners to start and stop")
	}
}
