package dockerStats

import (
	"context"
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pablogolobaro/dockertool-legend/internal/app"
	"go.uber.org/zap"
	"io"
	"os"
	"time"
)

type dockerStatsService struct {
	log      *zap.SugaredLogger
	cli      *client.Client
	shtdwnCh chan struct{}
}

func newDockerStatsService(log *zap.SugaredLogger, cli *client.Client) app.DockerService {
	shtdwn := make(chan struct{}, 1)
	return &dockerStatsService{log: log, cli: cli, shtdwnCh: shtdwn}
}

func (d *dockerStatsService) CollectStatsOnce() error {
	ctx := context.Background()
	containerList, err := d.cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err
	}

	for _, container := range containerList {
		containerStats, err := d.cli.ContainerStats(ctx, container.ID, false)
		if err != nil {
			return err
		}

		defer containerStats.Body.Close()

		io.Copy(os.Stdout, containerStats.Body)
	}

	return nil
}

func (d *dockerStatsService) StartStatsStream() error {
	ctx := context.Background()

	for {
		select {
		case <-d.shtdwnCh:
			return nil
		default:
			err := d.streamContainers(ctx)
			if err != nil {
				return err
			}
		}
	}

}

func (d *dockerStatsService) StopStatsStream() {
	d.shtdwnCh <- struct{}{}
}

func (d *dockerStatsService) streamContainers(ctx context.Context) error {
	containerList, err := d.cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err
	}

	tm.MoveCursor(1, 1)

	totals := tm.NewTable(0, 10, 5, ' ', 0)
	fmt.Fprintf(totals, "Container\tCPU\n")

	for _, container := range containerList {
		containerStats, err := d.cli.ContainerStats(ctx, container.ID, false)
		if err != nil {
			return err
		}

		defer containerStats.Body.Close()

		fmt.Fprintf(totals, "%s\t%d\n")

		io.Copy(os.Stdout, containerStats.Body)
	}

	tm.Flush() // Call it every time at the end of rendering

	time.Sleep(time.Second)

	return nil
}
