package gootstrap

// Runner defines a lifecycle factory that returns blocking start and stop
// functions for a long-lived process.
type Runner func() (start, stop func() error)
