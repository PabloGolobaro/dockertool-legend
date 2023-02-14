package dockerStats

import (
	"github.com/docker/docker/client"
	"github.com/pablogolobaro/dockertool-legend/internal/app"
	"github.com/pablogolobaro/dockertool-legend/internal/service/containerStreamer"
	"go.uber.org/zap"
)

type dockerStatsService struct {
	log      *zap.SugaredLogger
	cli      *client.Client
	shtdwnCh chan struct{}
	streamer *containerStreamer.ContainerStreamer
}

func NewDockerStatsService(shtdwn chan struct{}, log *zap.SugaredLogger, cli *client.Client, streamer *containerStreamer.ContainerStreamer) app.DockerService {
	return &dockerStatsService{log: log, cli: cli, shtdwnCh: shtdwn, streamer: streamer}
}
