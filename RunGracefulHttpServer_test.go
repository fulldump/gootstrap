package gootstrap

import (
	"net"
	"net/http"
	"testing"
	"time"
)

// TestRunGracefulHttpServer check that the HTTP server starts and stops correctly
func TestRunGracefulHttpServer(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to reserve free port: %v", err)
	}
	addr := ln.Addr().String()
	_ = ln.Close()

	s := &http.Server{
		Addr: addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("OK"))
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

	time.Sleep(200 * time.Millisecond)

	resp, err := http.Get("http://" + addr)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got: %v", resp.Status)
	}
	_ = resp.Body.Close()

	err = stop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	resp, err = http.Get("http://" + addr)
	if err == nil {
		_ = resp.Body.Close()
		t.Errorf("Expected error after server stop")
	}
}

// TestRunGracefulHttpServerWithGracefulShutdown check that the server correctly
// responds with a "Service Unavailable" status during the graceful shutdown period.
func TestRunGracefulHttpServerWithGracefulShutdown(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to reserve free port: %v", err)
	}
	addr := ln.Addr().String()
	_ = ln.Close()

	s := &http.Server{
		Addr: addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("OK"))
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

	time.Sleep(200 * time.Millisecond)

	go func() {
		err := stop()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	}()

	time.Sleep(1 * time.Second)

	resp, err := http.Get("http://" + addr)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("Expected status Service Unavailable, got: %v", resp.Status)
	}
	_ = resp.Body.Close()
}
