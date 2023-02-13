package dockerStats

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/pablogolobaro/dockertool-legend/internal/models"
	"io"
)

func (c *consoleWriter) getStats(ctx context.Context) error {
	containerList, err := c.cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err
	}

	statsList := make([]string, 0, len(containerList))

	for _, container := range containerList {
		containerStats, err := c.cli.ContainerStats(ctx, container.ID, false)
		if err != nil {
			return err
		}

		stat := models.Stats{Name: container.Image}
		bytes, err := io.ReadAll(containerStats.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(bytes, &stat)
		if err != nil {
			return err
		}

		oneTimeStat := fmt.Sprintf("%s\t%s\t%.2f\t%.2f\n", stat.Name, container.Image, stat.CalculateCPUUsage(), stat.CalculateMemoryUsage())

		statsList = append(statsList, oneTimeStat)

		containerStats.Body.Close()
	}

	writeOnce(statsList)

	return nil
}
