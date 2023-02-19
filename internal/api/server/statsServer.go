package server

import (
	pb "github.com/pablogolobaro/dockertool-legend/internal/api/protoStats"
	"github.com/pablogolobaro/dockertool-legend/internal/models"
	"go.uber.org/zap"
)

type ContainerApllication interface {
	CollectStatsOnce() ([]models.Stats, error)
	GetStreamChannel() (chan []models.Stats, chan struct{})
}

type StatsServer struct {
	pb.UnimplementedContainerStatsServiceServer
	App ContainerApllication
	Log *zap.SugaredLogger
}
