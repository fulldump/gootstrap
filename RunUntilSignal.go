package gootstrap

import (
	"log"
	"os"
	"os/signal"
	"sync"
)

// RunUntilSignal starts a runner and waits for specified signals to stop it.
func RunUntilSignal(run Runner, s ...os.Signal) {
	if run == nil {
		log.Println("ERROR: run: nil runner")
		return
	}

	start, stop := run()
	if start == nil || stop == nil {
		log.Println("ERROR: run: nil start or stop")
		return
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, s...)
	defer signal.Stop(sigs)

	done := make(chan struct{})
	stopOnce := sync.Once{}
	stopRunner := func() {
		stopOnce.Do(func() {
			if err := stop(); err != nil {
				log.Println("ERROR: stop:", err.Error())
			}
		})
	}

	go func() {
		select {
		case <-sigs:
			stopRunner()
		case <-done:
		}
	}()

	err := start()
	close(done)
	if err != nil {
		log.Println("ERROR: start:", err.Error())
	}
}
