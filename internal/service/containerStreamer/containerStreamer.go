package containerStreamer

import (
	"context"
	"github.com/docker/docker/client"
	"go.uber.org/zap"
	"sync"
	"time"
)

type ContainerStreamer struct {
	log                    *zap.SugaredLogger
	cli                    *client.Client
	shutdownCh             chan struct{}
	containerMap           map[string]struct{}
	containerStatsChannels []containerStatsChannel
	sync.WaitGroup
	sync.Mutex
}

type containerStatsChannel struct {
	statsCh chan string
	ID      string
}

func NewcontainerStreamer(log *zap.SugaredLogger, shutdownCh chan struct{}, cli *client.Client) *ContainerStreamer {
	containerMap := make(map[string]struct{})
	containerStatsChannels := make([]containerStatsChannel, 0)
	return &ContainerStreamer{log: log, shutdownCh: shutdownCh, cli: cli, containerMap: containerMap, containerStatsChannels: containerStatsChannels}
}

func (c *ContainerStreamer) StartStreaming(ctx context.Context, errCh chan error) chan []string {
	c.log.Debug("Start Streamer")

	streamChannel := make(chan []string)

	go func() {
		newContainers, delContainers := c.startNewContainersController(errCh)

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		c.Add(1)
		defer c.Done()
		defer close(delContainers)
		defer close(streamChannel)
		for {

			select {
			case <-ctx.Done():
				c.log.Debug("Stop Streamer")
				return

			case containerChan := <-newContainers:
				c.containerStatsChannels = append(c.containerStatsChannels, containerChan)

			case <-ticker.C:
				stats := make([]string, len(c.containerStatsChannels))

				for i, inContainer := range c.containerStatsChannels {
					stat, ok := <-inContainer.statsCh
					if !ok {
						c.containerStatsChannels[i] = c.containerStatsChannels[len(c.containerStatsChannels)-1]
						c.containerStatsChannels = c.containerStatsChannels[:len(c.containerStatsChannels)-1]

						delContainers <- inContainer.ID
					} else {
						stats = append(stats, stat)
					}
				}

				streamChannel <- stats
			}
		}
	}()
	return streamChannel
}
