package app

import (
	"context"
	"github.com/pablogolobaro/dockertool-legend/internal/config"
	"github.com/pablogolobaro/dockertool-legend/internal/models"
	"go.uber.org/zap"
)

type Apllication interface {
	Start(ctx context.Context)
	CollectStatsOnce() ([]models.Stats, error)
	GetStreamChannel() (chan []models.Stats, chan struct{})
	Error() chan error
}

type ContainerStreamer interface {
	StartStreaming(ctx context.Context, errCh chan error) chan []models.Stats
	GetStats(ctx context.Context) ([]models.Stats, error)
	WaitForAll()
}

type statsApp struct {
	log                  *zap.SugaredLogger
	mode                 *config.Mode
	errCh                chan error
	containerStreamer    ContainerStreamer
	streamsPool          []portStream
	newPortStreamChannel chan portStream
}

type portStream struct {
	statsCh  chan []models.Stats
	cancelCh chan struct{}
}
