package dockerStats

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/pablogolobaro/dockertool-legend/internal/models"
	"io"
	"os"
	"text/tabwriter"
)

func (d *dockerStatsService) getStats(ctx context.Context) error {
	containerList, err := d.cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return err
	}

	//tm.MoveCursor(1, 1)

	totals := tabwriter.NewWriter(os.Stdout, 0, 10, 5, ' ', 0)
	fmt.Fprintf(totals, "Container\tCPU %s\tMemory %s\n", "%", "%")

	for _, container := range containerList {
		containerStats, err := d.cli.ContainerStats(ctx, container.ID, false)
		if err != nil {
			return err
		}

		stats := models.Stats{}
		bytes, err := io.ReadAll(containerStats.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(bytes, &stats)
		if err != nil {
			return err
		}

		fmt.Fprintf(totals, "%s\t%.2f\t%.2f\n", container.Image, stats.CalculateCPUUsage(), stats.CalculateMemoryUsage())

		containerStats.Body.Close()
	}

	totals.Flush()

	// Call it every time at the end of rendering

	return nil
}
