package containerstreamer

import (
	"context"
	"github.com/pablogolobaro/dockertool-legend/internal/app"
	"github.com/pablogolobaro/dockertool-legend/internal/models"
	"time"
)

type mockStreamer struct {
}

func NewMockStreamer() app.ContainerStreamer {
	return &mockStreamer{}
}

func (m *mockStreamer) StartStreaming(ctx context.Context, errCh chan error) chan []models.Stats {
	c := make(chan []models.Stats)
	go func() {
		for {
			c <- []models.Stats{{ID: "11111111111", Image: "Postgres", Name: "Panda", MemoryStats: struct {
				MaxUsage int `json:"max_usage"`
				Usage    int `json:"usage"`
				Limit    int `json:"limit"`
			}{
				MaxUsage: 10,
				Usage:    2,
				Limit:    20,
			}, CpuStats: struct {
				CpuUsage struct {
					TotalUsage int `json:"total_usage"`
				} `json:"cpu_usage"`
				SystemCpuUsage int64 `json:"system_cpu_usage"`
			}{CpuUsage: struct {
				TotalUsage int `json:"total_usage"`
			}{
				TotalUsage: 10,
			}, SystemCpuUsage: 100}}}

			time.Sleep(time.Second)
		}
	}()
	return c
}

func (m *mockStreamer) GetStats(ctx context.Context) ([]models.Stats, error) {
	stats := []models.Stats{{ID: "11111111111", Image: "Postgres", Name: "Panda", MemoryStats: struct {
		MaxUsage int `json:"max_usage"`
		Usage    int `json:"usage"`
		Limit    int `json:"limit"`
	}{
		MaxUsage: 10,
		Usage:    2,
		Limit:    20,
	}, CpuStats: struct {
		CpuUsage struct {
			TotalUsage int `json:"total_usage"`
		} `json:"cpu_usage"`
		SystemCpuUsage int64 `json:"system_cpu_usage"`
	}{CpuUsage: struct {
		TotalUsage int `json:"total_usage"`
	}{
		TotalUsage: 10,
	}, SystemCpuUsage: 100}}}

	return stats, nil
}

func (m *mockStreamer) WaitForAll() {
	time.Sleep(time.Second)
}
