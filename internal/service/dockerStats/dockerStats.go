package dockerStats

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pablogolobaro/dockertool-legend/internal/app"
	"go.uber.org/zap"
)

type dockerStatsService struct {
	log           *zap.SugaredLogger
	cli           *client.Client
	shtdwnCh      chan struct{}
	consoleWriter *consoleWriter
}

func NewDockerStatsService(shtdwn chan struct{}, log *zap.SugaredLogger, cli *client.Client, consoleWriter *consoleWriter) app.DockerService {
	return &dockerStatsService{log: log, cli: cli, shtdwnCh: shtdwn, consoleWriter: consoleWriter}
}

func (d *dockerStatsService) CollectStatsOnce() error {
	ctx := context.Background()
	return d.getStats(ctx)
}

func (d *dockerStatsService) StartStatsStream() error {
	ctx := context.Background()
	containerList, err := d.cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err
	}
	d.consoleWriter.Add(1)
	go d.consoleWriter.StartWriteToConsole()
	for _, container := range containerList {
		d.consoleWriter.Add(1)
		go d.consoleWriter.streamContainer(ctx, container)
	}
	for {
		select {
		case <-d.shtdwnCh:
			return nil
		}
	}
}

func (d *dockerStatsService) StopStatsStream() {
	close(d.shtdwnCh)
	d.consoleWriter.Wait()
}
