// Package gootstrap provides small, composable helpers to bootstrap long-lived
// Go services.
//
// A Runner returns two blocking lifecycle functions:
//   - start, which starts and blocks while the service is running.
//   - stop, which gracefully stops the service.
//
// Use Run for the common case (start and stop on SIGINT/SIGTERM), or combine
// multiple runners with RunAll.
package gootstrap
