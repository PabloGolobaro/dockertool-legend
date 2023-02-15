package config

import (
	"flag"
)

type Flags struct {
	Console bool
	Port    int
}

func GetModeFlags() (Flags, error) {
	console := flag.Bool("console", false, "Output to StdOut")
	port := flag.Int("port", 50051, "gRPC server port")

	flag.Parse()

	return Flags{Console: *console, Port: *port}, nil
}
