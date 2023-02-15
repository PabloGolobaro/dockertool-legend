package containerstreamer

import (
	"bufio"
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/pablogolobaro/dockertool-legend/internal/models"
)

func (c *ContainerStreamer) streamContainer(ctx context.Context, container types.Container) chan models.Stats {
	containerChan := make(chan models.Stats)

	c.Add(1)

	go func() {
		defer c.Done()
		defer close(containerChan)

		containerStats, err := c.cli.ContainerStats(ctx, container.ID, true)
		if err != nil {
			c.log.Errorw("Stream container error", "error", err, "container", container.ID)

			return
		}

		c.log.Debugw("Stream", "container", container.ID)

		defer containerStats.Body.Close()

		scanner := bufio.NewScanner(containerStats.Body)
		for scanner.Scan() {
			stats := models.Stats{ID: container.ID, Name: container.Image, Image: container.Image}

			err = json.Unmarshal(scanner.Bytes(), &stats)
			if err != nil {
				c.log.Errorw("Stream container error", "error", err, "container", container.ID)

				return
			}
			select {
			case <-ctx.Done():
				c.log.Debugw("Stream closed", "container", container.ID)

				return
			case containerChan <- stats:

			}
		}
		c.log.Debugw("Stream closed", "container", container.ID)

	}()
	return containerChan
}
