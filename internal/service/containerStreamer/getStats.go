package containerStreamer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/pablogolobaro/dockertool-legend/internal/models"
	"io"
)

func (c *ContainerStreamer) GetStats(ctx context.Context) ([]string, error) {
	containerList, err := c.cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	statsList := make([]string, 0, len(containerList))

	for _, container := range containerList {
		containerStats, err := c.cli.ContainerStats(ctx, container.ID, false)
		if err != nil {
			return statsList, err
		}

		stat := models.Stats{Name: container.Image}
		bytes, err := io.ReadAll(containerStats.Body)
		if err != nil {
			return statsList, err
		}

		err = json.Unmarshal(bytes, &stat)
		if err != nil {
			return statsList, err
		}

		oneTimeStat := fmt.Sprintf("%s\t%s\t%.2f\t%.2f\n", stat.Name, container.Image, stat.CalculateCPUUsage(), stat.CalculateMemoryUsage())

		statsList = append(statsList, oneTimeStat)

		containerStats.Body.Close()
	}

	return statsList, nil
}
