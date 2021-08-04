// Package cmd provides the command line interface (cli).
package cmd

import "time"

const (
	// DefautAddress is the default HTTP address
	DefautAddress = ":8080"
	// DefaultShutdownTimeout is the default shutdown timeout. Shutdown timeout before connections are cancelled
	DefaultShutdownTimeout = 10 * time.Second
	// DefaultReadTimeout is the default read timeout before connection is cancelled
	DefaultReadTimeout = 5 * time.Second
	// DefaultWriteTimeout is the default write timeout before connection is cancelled
	DefaultWriteTimeout = 10 * time.Second
	// DefaultCors is the default Cross Origin Resource Sharing AllowedOrigins
	DefaultCors = "*"
)
