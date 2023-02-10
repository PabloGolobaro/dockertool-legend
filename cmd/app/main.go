package main

import (
	"github.com/pablogolobaro/dockertool-legend/internal/app"
	"github.com/pablogolobaro/dockertool-legend/internal/config"
	"github.com/pablogolobaro/dockertool-legend/internal/logger"
	"github.com/pablogolobaro/dockertool-legend/internal/service/dockerStats"
	"github.com/pablogolobaro/dockertool-legend/pkg/docker"
	"go.uber.org/zap"
	"sync"
)

var (
	log           *zap.SugaredLogger
	mode          *app.Mode
	dockerService app.DockerService
	builder       app.AppBuilderInt
)

func main() {

	log.Info("Building Application")
	application := builder.Logger(log).Mode(mode).DockerService(dockerService).Build()

	var wg sync.WaitGroup
	log.Info("Start Application")
	wg.Add(1)
	go func(application app.Apllication) {
		application.Start()

		defer wg.Done()
	}(application)

	wg.Wait()
	log.Info("Gracefully Stopped Application")

}

func init() {
	var err error
	log, err = logger.NewLogger()
	if err != nil {
		panic(err)
	}

	log.Debug("Creating DockerClient")

	log.Debug("Get flags")
	modeFlags, err := config.GetModeFlags()
	if err != nil {
		log.Fatal(err)
	}

	log.Debugw("Got flags", "Flags", modeFlags)

	mode = app.NewMode(modeFlags.Stream, modeFlags.Duration)

	dockerClient, err := docker.NewDockerClient()
	if err != nil {
		log.Fatal(err)
	}

	shtdwn := make(chan struct{}, 1)

	consoleWriter := dockerStats.NewConsoleWriter(log, shtdwn, dockerClient)

	dockerService = dockerStats.NewDockerStatsService(shtdwn, log, dockerClient, consoleWriter)

	builder = app.NewDockerAppBuilder()
}
