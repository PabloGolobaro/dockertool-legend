package server

import (
	pb "github.com/pablogolobaro/dockertool-legend/internal/api/protoStats"
	"github.com/pablogolobaro/dockertool-legend/internal/app"
)

type StatsServer struct {
	pb.UnimplementedContainerStatsServiceServer
	dockerService app.DockerService
}

func (s StatsServer) GetStatsStream(message *pb.GetStatsMessage, stream pb.ContainerStatsService_GetStatsStreamServer) error {
	select {
	case <-stream.Context().Done():
		return nil

	}

}
