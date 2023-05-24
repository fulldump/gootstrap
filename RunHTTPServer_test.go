package gootstrap

import (
	"net/http"
	"testing"
	"time"
)

func TestRunHTTPServer(t *testing.T) {
	server := &http.Server{
		Addr: "127.0.0.1:8080", // use a random free port
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("Hello, world!"))
		}),
	}

	runner := RunHTTPServer(server)
	start, stop := runner()

	errCh := make(chan error)
	go func() {
		errCh <- start()
	}()

	// Give server some time to start
	time.Sleep(1 * time.Second)

	// Send request to the server
	resp, err := http.Get("http://" + server.Addr)
	if err != nil {
		t.Fatalf("Failed to send request to server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}

	// Check if the server returns error ErrServerClosed after being stopped
	err = stop()
	if err != nil {
		t.Fatalf("Failed to stop server: %v", err)
	}

	err = <-errCh
	if err != http.ErrServerClosed {
		t.Errorf("Unexpected error: got %v, want %v", err, http.ErrServerClosed)
	}
}
