package dockerStats

import (
	"bufio"
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/pablogolobaro/dockertool-legend/internal/models"
	"time"
)

func (c *consoleWriter) streamContainer(ctx context.Context, container types.Container) {
	containerStats, err := c.cli.ContainerStats(ctx, container.ID, true)
	if err != nil {
		c.log.Errorw("Stream container error", "error", err, "container", container.Image)
	}
	c.log.Debugw("Stream", "container", container.Image)
	defer containerStats.Body.Close()
	scanner := bufio.NewScanner(containerStats.Body)
	for scanner.Scan() {
		select {
		case <-c.shtdwn:
			c.log.Debugw("Stream closed", "container", container.Image)
			c.Done()
			return
		default:
			stats := models.Stats{Name: container.Image}

			err = json.Unmarshal(scanner.Bytes(), &stats)
			if err != nil {
				c.log.Errorw("Stream container error", "error", err, "container", container.Image)
			}

			//c.log.Debugw("Stream", "container", container.Image, "stats", stats)

			c.input <- stats
			time.Sleep(time.Millisecond * 10)
		}
	}
}
