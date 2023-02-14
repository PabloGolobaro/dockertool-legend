package dockerStats

import (
	"context"
	"github.com/pablogolobaro/dockertool-legend/pkg/console"
)

func (d *dockerStatsService) CollectStatsOnce() error {
	ctx := context.Background()
	stats, err := d.streamer.GetStats(ctx)
	if err != nil {
		return err
	}
	console.WriteOnce(stats)
	return nil
}
