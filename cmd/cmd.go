// Package cmd provides the command line interface (cli).
package cmd

import "time"

const (
	DefautAddress          = ":8080"
	DefaultShutdownTimeout = 10 * time.Second
	DefaultReadTimeout     = 5 * time.Second
	DefaultWriteTimeout    = 10 * time.Second
	DefaultCors            = "*"
)
