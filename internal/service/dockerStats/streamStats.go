package dockerStats

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/pablogolobaro/dockertool-legend/internal/models"
)

func (c *consoleWriter) streamContainer(ctx context.Context, container types.Container) chan string {
	containerChan := make(chan string)

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
			stats := models.Stats{Name: container.Image}

			err = json.Unmarshal(scanner.Bytes(), &stats)
			if err != nil {
				c.log.Errorw("Stream container error", "error", err, "container", container.ID)

				return
			}

			oneTimeStat := fmt.Sprintf("%s\t%s\t%.2f\t%.2f\n", stats.Name, container.Image, stats.CalculateCPUUsage(), stats.CalculateMemoryUsage())

			select {
			case <-ctx.Done():
				c.log.Debugw("Stream closed", "container", container.ID)

				return
			case containerChan <- oneTimeStat:

			}
		}
		c.log.Debugw("Stream closed", "container", container.ID)

	}()
	return containerChan
}
