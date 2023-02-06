package app

import (
	"go.uber.org/zap"
)

type Apllication interface {
	Start()
}

type DockerService interface {
	CollectStatsOnce() error
	StartStatsStream() error
	StopStatsStream()
}

type dockerStatsApp struct {
	log           *zap.SugaredLogger
	mode          Mode
	errCh         chan error
	dockerService DockerService
}

func (d *dockerStatsApp) Start() {
	d.log.Info("App starts collecting stats")
	if !d.mode.stream {
		err := d.dockerService.CollectStatsOnce()
		if err != nil {
			d.log.Errorw("error collecting stats one time", "Error", err)
		}
		return
	}

	go func() {
		err := d.dockerService.StartStatsStream()
		if err != nil {
			d.errCh <- err
		}
	}()

	if d.mode.WithTimer {
		d.WaitWithTimer()
	} else {
		d.WaitWithoutTimer()
	}

	d.log.Info("App ends collecting stats")
}