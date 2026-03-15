package gootstrap

import (
	"context"
	"fmt"
	"net/http"
)

// RunHTTPServer bootstraps a standard net/http server lifecycle.
func RunHTTPServer(server *http.Server) Runner {
	return func() (start, stop func() error) {

		start = func() error {
			fmt.Println("HTTP server listening on", server.Addr)
			return server.ListenAndServe()
		}

		stop = func() error {
			return server.Shutdown(context.Background())
		}

		return
	}
}
