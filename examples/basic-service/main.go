package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fulldump/gootstrap"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("ok"))
		}),
	}

	runner := gootstrap.RunAll(
		gootstrap.RunHTTPServer(server),
		newTickerWorker("billing-sync", 2*time.Second),
	)

	gootstrap.RunUntilSignal(runner)
}

func newTickerWorker(name string, every time.Duration) gootstrap.Runner {
	return func() (start, stop func() error) {
		stopCh := make(chan struct{})
		stopOnce := sync.Once{}

		start = func() error {
			t := time.NewTicker(every)
			defer t.Stop()

			for {
				select {
				case <-t.C:
					log.Printf("worker[%s]: tick", name)
				case <-stopCh:
					fmt.Printf("worker[%s]: stopped\n", name)
					return nil
				}
			}
		}

		stop = func() error {
			stopOnce.Do(func() {
				close(stopCh)
			})
			return nil
		}

		return start, stop
	}
}
