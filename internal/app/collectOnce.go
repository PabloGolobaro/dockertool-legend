package app

import (
	"context"
	"github.com/pablogolobaro/dockertool-legend/internal/models"
)

func (a *statsApp) CollectStatsOnce() ([]models.Stats, error) {
	ctx := context.Background()
	stats, err := a.containerStreamer.GetStats(ctx)
	if err != nil {
		return stats, err
	}
	return stats, nil
}
