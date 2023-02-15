package main

import "flag"

type Flags struct {
	Timeout int
	Port    int
}

func getFlags() (Flags, error) {
	timeout := flag.Int("timeout", 0, "Timeout")
	port := flag.Int("port", 50051, "gRPC server port")

	flag.Parse()

	return Flags{Timeout: *timeout, Port: *port}, nil
}
