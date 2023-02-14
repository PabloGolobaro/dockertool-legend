package dockerStats

import (
	"context"
	"github.com/pablogolobaro/dockertool-legend/pkg/console"
)

func (d *dockerStatsService) StartStatsStream(errCh chan error) {
	cwc, cancel := context.WithCancel(context.Background())

	streamChannel := d.streamer.StartStreaming(cwc, errCh)

	for {
		select {
		case stats := <-streamChannel:
			console.WriteToConsole(stats)
		case <-d.shtdwnCh:
			cancel()
			return
		}
	}
}

func (d *dockerStatsService) StopStatsStream() {
	close(d.shtdwnCh)
	d.streamer.Wait()
}
