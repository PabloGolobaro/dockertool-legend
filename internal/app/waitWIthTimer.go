package app

import (
	"os"
	"os/signal"
	"time"
)

func (d *dockerStatsApp) WaitWithTimer() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	timer := time.NewTimer(d.mode.duration)

	select {
	case <-timer.C:
		d.log.Info("Stop collecting stats by timer")
		d.dockerService.StopStatsStream()
	case <-stop:
		d.log.Info("Stop collecting stats")
		d.dockerService.StopStatsStream()
	case err := <-d.errCh:
		d.log.Errorw("Stop collecting stats", "Error", err)
	}
}

func (d *dockerStatsApp) WaitWithoutTimer() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	select {
	case <-stop:
		d.log.Debug("Stop collecting stats from wait")
		d.dockerService.StopStatsStream()
	case err := <-d.errCh:
		d.log.Errorw("Stop collecting stats", "Error", err)
	}
}
