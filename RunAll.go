package gootstrap

import (
	"log"
	"sync"
)

// RunAll accepts a list of runners and returns a new runner that starts all of
// them on start and stops all of them on stop.
func RunAll(runners ...Runner) Runner {
	return func() (start func() error, stop func() error) {

		wgStart := sync.WaitGroup{}
		stops := []func() error{}
		stopsMu := sync.Mutex{}
		stopOnce := sync.Once{}

		start = func() error {

			for i, run := range runners {
				if run == nil {
					log.Printf("runall: runner[%d] is nil, skipping", i)
					continue
				}

				runnerStart, runnerStop := run()
				if runnerStart == nil || runnerStop == nil {
					log.Printf("runall: runner[%d] returned nil start or stop, skipping", i)
					continue
				}

				stopsMu.Lock()
				stops = append(stops, runnerStop)
				stopsMu.Unlock()
				wgStart.Add(1)
				go func(start func() error) {
					// TODO: handle panics
					defer wgStart.Done()
					err := start() // should be a blocking call
					if err != nil {
						log.Println("start:", err.Error())
					}
				}(runnerStart)
			}

			wgStart.Wait()

			return nil
		}

		wgStop := sync.WaitGroup{}
		stop = func() error {
			stopOnce.Do(func() {
				stopsMu.Lock()
				stopsCopy := append([]func() error(nil), stops...)
				stopsMu.Unlock()

				for _, stop := range stopsCopy {
					wgStop.Add(1)
					go func(stop func() error) {
						defer wgStop.Done()
						err := stop()
						if err != nil {
							log.Println("stop:", err.Error())
						}
					}(stop)
				}

				wgStop.Wait()
			})

			return nil
		}

		return start, stop
	}
}
