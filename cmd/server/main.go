package main

import (
	"context"
	pb "github.com/pablogolobaro/dockertool-legend/internal/api/protoStats"
	"github.com/pablogolobaro/dockertool-legend/internal/api/server"
	"github.com/pablogolobaro/dockertool-legend/internal/app"
	"github.com/pablogolobaro/dockertool-legend/internal/config"
	"github.com/pablogolobaro/dockertool-legend/internal/logger"
	"github.com/pablogolobaro/dockertool-legend/internal/service/containerstreamer"
	"github.com/pablogolobaro/dockertool-legend/pkg/docker"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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

	grpcServer = grpc.NewServer(grpc.StreamInterceptor(server.StatsServerStreamInterceptor))

	pb.RegisterContainerStatsServiceServer(grpcServer, &server.StatsServer{App: application, Log: log})

	var wg sync.WaitGroup
	go func() {
		wg.Add(1)

		log.Info("Start Application")
		application.Start(cwc)

		defer wg.Done()
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
