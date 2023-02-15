package main

import (
	"google.golang.org/grpc"
	"net"
)

func startgRPCServer(lis net.Listener) {
	log.Info("Starting gRPC listener on port " + mode.Port())
	if err := grpcServer.Serve(lis); err != nil {
		if err == grpc.ErrServerStopped {
			log.Info("Stopping gRPC server")
		} else {
			log.Fatalf("failed to serve: %v", err)
		}
	}
}
