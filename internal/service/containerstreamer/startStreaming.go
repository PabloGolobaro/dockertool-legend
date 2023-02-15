package containerstreamer

import (
	"context"
	"github.com/pablogolobaro/dockertool-legend/internal/models"
	"time"
)

func (c *ContainerStreamer) StartStreaming(ctx context.Context, errCh chan error) chan []models.Stats {
	c.log.Debug("Start Streamer")

	streamChannel := make(chan []models.Stats)

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
				stats := make([]models.Stats, len(c.containerStatsChannels))

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
