package dockerStats

import (
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/docker/docker/client"
	"github.com/pablogolobaro/dockertool-legend/internal/models"
	"go.uber.org/zap"
	"sync"
	"time"
)

type consoleWriter struct {
	log      *zap.SugaredLogger
	input    chan models.Stats
	shtdwn   chan struct{}
	statsMap map[string]models.Stats
	sync.WaitGroup
	cli *client.Client
}

func NewConsoleWriter(log *zap.SugaredLogger, shtdwn chan struct{}, cli *client.Client) *consoleWriter {
	statsMap := make(map[string]models.Stats)

	input := make(chan models.Stats, 10)

	return &consoleWriter{log: log, shtdwn: shtdwn, cli: cli, statsMap: statsMap, input: input}
}

func (c *consoleWriter) StartWriteToConsole() {
	c.log.Debug("Start Console Writer")
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {

		select {
		case <-ticker.C:

			tm.Clear()

			table := tm.NewTable(0, 10, 5, ' ', 0)

			fmt.Fprintf(table, "Container\tCPU %s\tMemory %s\n", "%", "%")

			//c.log.Debugw("Writer reads map", "map", c.statsMap)
			for _, stats := range c.statsMap {
				fmt.Fprintf(table, "%s\t%.2f\t%.2f\n", stats.Name, stats.CalculateCPUUsage(), stats.CalculateMemoryUsage())

			}
			//tm.MoveCursorUp(7)

			tm.Print(table)

			tm.Flush()

		case stats := <-c.input:
			//c.log.Debugw("Writer input chanel", "stats", stats)
			c.statsMap[stats.Name] = stats
		case <-c.shtdwn:
			c.Done()
			return
		}

	}
}
