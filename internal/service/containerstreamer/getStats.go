package containerstreamer

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/pablogolobaro/dockertool-legend/internal/models"
	"io"
)

func (c *ContainerStreamer) GetStats(ctx context.Context) ([]models.Stats, error) {
	containerList, err := c.cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	statsList := make([]models.Stats, 0, len(containerList))

	for _, container := range containerList {
		containerStats, err := c.cli.ContainerStats(ctx, container.ID, false)
		if err != nil {
			return statsList, err
		}

		stat := models.Stats{ID: container.ID, Name: container.Image, Image: container.Image}
		bytes, err := io.ReadAll(containerStats.Body)
		if err != nil {
			return statsList, err
		}

		err = json.Unmarshal(bytes, &stat)
		if err != nil {
			return statsList, err
		}

		statsList = append(statsList, stat)

		containerStats.Body.Close()
	}

	return statsList, nil
}
