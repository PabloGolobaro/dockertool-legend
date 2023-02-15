package console

import (
	"github.com/pablogolobaro/dockertool-legend/internal/models"
	"testing"
)

func TestWriteToConsole(t *testing.T) {
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

	tests := []struct {
		name  string
		stats []models.Stats
	}{
		{name: "First", stats: stats},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < 100; i++ {
				WriteToConsole(tt.stats)
			}
		})
	}
}
