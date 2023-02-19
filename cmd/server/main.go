package main

import (
	"context"
	"crypto/tls"
	pb "github.com/pablogolobaro/dockertool-legend/internal/api/protoStats"
	"github.com/pablogolobaro/dockertool-legend/internal/api/server"
	"github.com/pablogolobaro/dockertool-legend/internal/app"
	"github.com/pablogolobaro/dockertool-legend/internal/config"
	"github.com/pablogolobaro/dockertool-legend/internal/logger"
	"github.com/pablogolobaro/dockertool-legend/internal/service/containerstreamer"
	"github.com/pablogolobaro/dockertool-legend/pkg/docker"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"os"
	"os/signal"
	"sync"
)

var (
	log               *zap.SugaredLogger
	mode              *config.Mode
	builder           app.AppBuilderInt
	containerStreamer app.ContainerStreamer
	grpcServer        *grpc.Server
)

const (
	crtFile = "./certs/server.crt"
	keyFile = "./certs/server.key"
)

func main() {
	cwc, cancel := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	log.Info("Building Application")

	application := builder.Logger(log).Mode(mode).ContainerStreamer(containerStreamer).Build()

	lis, err := net.Listen("tcp", "localhost:"+mode.Port())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	certificate, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load key pair: %v", err)
	}

	options := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&certificate)),
		grpc.UnaryInterceptor(server.EnsureValidBasicCredentials),
		grpc.StreamInterceptor(server.StatsServerStreamInterceptor),
	}

	grpcServer = grpc.NewServer(options...)

	pb.RegisterContainerStatsServiceServer(grpcServer, &server.StatsServer{App: application, Log: log})

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		log.Info("Start Application")
		application.Start(cwc)

		wg.Done()
	}()

	go startgRPCServer(lis)

	select {
	case <-sigCh:
		log.Info("Stop signal received")
	case err := <-application.Error():
		log.Errorw("Stop collecting stats", "Error", err)
	}

	grpcServer.GracefulStop()

	cancel()

	wg.Wait()

	log.Info("Gracefully Stopped Application")

}

func init() {
	var err error
	log, err = logger.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("Creating DockerClient")

	log.Debug("Get flags")
	modeFlags, err := config.GetModeFlags()
	if err != nil {
		log.Fatal(err)
	}

	log.Debugw("Got flags", "Flags", modeFlags)

	mode = config.NewMode(modeFlags.Console, modeFlags.Port)

	dockerClient, err := docker.NewDockerClient()
	if err != nil {
		log.Fatal(err)
	}

	containerStreamer = containerstreamer.NewContainerStreamer(log, dockerClient)

	//containerStreamer = containerstreamer.NewMockStreamer()

	builder = app.NewDockerAppBuilder()

}
