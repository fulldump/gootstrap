package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fulldump/gootstrap"
)

func main() {
	var draining atomic.Bool
	var jobsProcessed atomic.Int64

	apiServer := &http.Server{
		Addr:    ":8080",
		Handler: apiHandler(&draining),
	}

	metricsServer := &http.Server{
		Addr:    ":9090",
		Handler: metricsHandler(&jobsProcessed, &draining),
	}

	runner := gootstrap.RunAll(
		withDrainingFlag(gootstrap.RunGracefulHttpServer(apiServer), &draining),
		gootstrap.RunHTTPServer(metricsServer),
		newWorker("invoice-sync", 2*time.Second, &jobsProcessed),
	)

	gootstrap.RunUntilSignal(runner)
}

func apiHandler(draining *atomic.Bool) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello from api\n"))
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})

	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if draining.Load() {
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte("draining\n"))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ready\n"))
	})

	return mux
}

func metricsHandler(processed *atomic.Int64, draining *atomic.Bool) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("jobs_processed " + strconv.FormatInt(processed.Load(), 10) + "\n"))
		if draining.Load() {
			_, _ = w.Write([]byte("service_draining 1\n"))
			return
		}
		_, _ = w.Write([]byte("service_draining 0\n"))
	})

	return mux
}

func newWorker(name string, every time.Duration, processed *atomic.Int64) gootstrap.Runner {
	return func() (start, stop func() error) {
		stopCh := make(chan struct{})
		stopOnce := sync.Once{}

		start = func() error {
			t := time.NewTicker(every)
			defer t.Stop()

			for {
				select {
				case <-t.C:
					n := processed.Add(1)
					log.Printf("worker[%s]: processed batch %d", name, n)
				case <-stopCh:
					log.Printf("worker[%s]: stopped", name)
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

func withDrainingFlag(runner gootstrap.Runner, draining *atomic.Bool) gootstrap.Runner {
	return func() (func() error, func() error) {
		start, stop := runner()

		wrappedStop := func() error {
			draining.Store(true)
			fmt.Println("service entering drain mode")
			return stop()
		}

		return start, wrappedStop
	}
}
