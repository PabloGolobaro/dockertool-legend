package models

type Stats struct {
	Name        string `json:"name,omitempty"`
	MemoryStats struct {
		MaxUsage int `json:"max_usage"`
		Usage    int `json:"usage"`
		Limit    int `json:"limit"`
	} `json:"memory_stats"`
	CpuStats struct {
		CpuUsage struct {
			TotalUsage int `json:"total_usage"`
		} `json:"cpu_usage"`
		SystemCpuUsage int64 `json:"system_cpu_usage"`
	} `json:"cpu_stats"`
}

func (s *Stats) CalculateCPUUsage() float64 {
	return float64(s.CpuStats.CpuUsage.TotalUsage) / float64(int(s.CpuStats.SystemCpuUsage)) * 100
}

func (s *Stats) CalculateMemoryUsage() float64 {
	return float64(s.MemoryStats.Usage) / float64(s.MemoryStats.MaxUsage) * 100
}
