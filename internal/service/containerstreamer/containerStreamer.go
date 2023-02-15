package containerstreamer

import (
	"github.com/docker/docker/client"
	"github.com/pablogolobaro/dockertool-legend/internal/app"
	"github.com/pablogolobaro/dockertool-legend/internal/models"
	"go.uber.org/zap"
	"sync"
)

type ContainerStreamer struct {
	log                    *zap.SugaredLogger
	cli                    *client.Client
	containerMap           map[string]struct{}
	containerStatsChannels []containerStatsChannel
	sync.WaitGroup
	sync.Mutex
}

type containerStatsChannel struct {
	statsCh chan models.Stats
	ID      string
}

func NewContainerStreamer(log *zap.SugaredLogger, cli *client.Client) app.ContainerStreamer {
	containerMap := make(map[string]struct{})
	containerStatsChannels := make([]containerStatsChannel, 0)
	return &ContainerStreamer{log: log, cli: cli, containerMap: containerMap, containerStatsChannels: containerStatsChannels}
}
