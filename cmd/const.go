package cmd

import "time"

const (
	DefautAddress          = ":13000"
	DefaultShutdownTimeout = 10 * time.Second
	DefaultReadTimeout     = 5 * time.Second
	DefaultWriteTimeout    = 10 * time.Second
	DefaultCors            = "*"
)
