package dockerStats

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/pablogolobaro/dockertool-legend/internal/app"
	"go.uber.org/zap"
	"time"
)

type dockerStatsService struct {
	log      *zap.SugaredLogger
	cli      *client.Client
	shtdwnCh chan struct{}
}

func NewDockerStatsService(log *zap.SugaredLogger, cli *client.Client) app.DockerService {
	shtdwn := make(chan struct{}, 1)
	return &dockerStatsService{log: log, cli: cli, shtdwnCh: shtdwn}
}

func (d *dockerStatsService) CollectStatsOnce() error {
	ctx := context.Background()
	return d.getStats(ctx)
}

func (d *dockerStatsService) StartStatsStream() error {
	ctx := context.Background()

	for {
		select {
		case <-d.shtdwnCh:
			return nil
		default:
			err := d.getStats(ctx)
			if err != nil {
				return err
			}
			time.Sleep(time.Second)
		}
	}

}

func (d *dockerStatsService) StopStatsStream() {
	d.shtdwnCh <- struct{}{}
}
