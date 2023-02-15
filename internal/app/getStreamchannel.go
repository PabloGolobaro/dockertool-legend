package app

import "github.com/pablogolobaro/dockertool-legend/internal/models"

func (a *statsApp) GetStreamChannel() (chan []models.Stats, chan struct{}) {
	ch := make(chan []models.Stats, 0)
	cancel := make(chan struct{}, 0)

	a.newPortStreamChannel <- portStream{statsCh: ch, cancelCh: cancel}

	return ch, cancel
}
