package dockerStats

import (
	"context"
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
	return d.consoleWriter.getStats(ctx)
}

func (d *dockerStatsService) StartStatsStream(errCh chan error) error {

	go d.consoleWriter.startWriteToConsole(errCh)

	<-d.shtdwnCh

	return nil
}

func (d *dockerStatsService) StopStatsStream() {
	close(d.shtdwnCh)
	d.consoleWriter.Wait()
}
