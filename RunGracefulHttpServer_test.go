package gootstrap

import (
	"net/http"
	"testing"
	"time"
)

// TestRunGracefulHttpServer check that the HTTP server starts and stops correctly
func TestRunGracefulHttpServer(t *testing.T) {
	s := &http.Server{
		Addr: ":8081",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		}),
	}

	runner := RunGracefulHttpServer(s)

	start, stop := runner()

	go func() {
		err := start()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}()

	resp, err := http.Get("http://localhost:8081")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got: %v", resp.Status)
	}

	err = stop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	resp, err = http.Get("http://localhost:8081")
	if err == nil {
		t.Errorf("Expected error after server stop")
	}
}

// TestRunGracefulHttpServerWithGracefulShutdown check that the server correctly
// responds with a "Service Unavailable" status during the graceful shutdown period.
func TestRunGracefulHttpServerWithGracefulShutdown(t *testing.T) {
	s := &http.Server{
		Addr: ":8082",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		}),
	}

	runner := RunGracefulHttpServer(s)

	start, stop := runner()

	go func() {
		err := start()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}()

	time.Sleep(1 * time.Second)

	go func() {
		err := stop()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}()

	time.Sleep(1 * time.Second)

	resp, err := http.Get("http://localhost:8082")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("Expected status Service Unavailable, got: %v", resp.Status)
	}
}
