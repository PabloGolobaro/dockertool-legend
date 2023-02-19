package server

import pb "github.com/pablogolobaro/dockertool-legend/internal/api/protoStats"

func (s *StatsServer) GetStatsStream(message *pb.GetStatsMessage, stream pb.ContainerStatsService_GetStatsStreamServer) error {
	ctx := stream.Context()
	channel, cancel := s.App.GetStreamChannel()
	for {
		select {
		case <-ctx.Done():
			s.Log.Debug(ctx.Err())
			close(cancel)
			return nil
		case stats := <-channel:

			resultStats := &pb.Stats{}

			for _, stat := range stats {
				containerStats := &pb.ContainerStats{
					Container: &pb.Container{Image: stat.Image, Name: stat.Name, Id: stat.ID},
					CPU:       float32(stat.CalculateCPUUsage()),
					Memory:    float32(stat.CalculateMemoryUsage()),
				}
				resultStats.ContainerStats = append(resultStats.GetContainerStats(), containerStats)
			}

			stream.Send(resultStats)
		}
	}
}
