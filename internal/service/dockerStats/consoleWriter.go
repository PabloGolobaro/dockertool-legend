package dockerStats

import (
	"github.com/docker/docker/client"
	"go.uber.org/zap"
	"sync"
	"time"
)

type consoleWriter struct {
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

func NewConsoleWriter(log *zap.SugaredLogger, shutdownCh chan struct{}, cli *client.Client) *consoleWriter {
	containerMap := make(map[string]struct{})
	containerStatsChannels := make([]containerStatsChannel, 0)
	return &consoleWriter{log: log, shutdownCh: shutdownCh, cli: cli, containerMap: containerMap, containerStatsChannels: containerStatsChannels}
}

func (c *consoleWriter) startWriteToConsole(errCh chan error) {
	c.log.Debug("Start Console Writer")

	newContainers, delContainers := c.startNewContainersController(errCh)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	c.Add(1)
	defer c.Done()
	defer close(delContainers)
	for {

		select {
		case <-c.shutdownCh:
			c.log.Debug("Stop Console Writer")
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

			writeToConsole(stats)
		}
	}
}
