package console

import (
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/pablogolobaro/dockertool-legend/internal/models"
	"os"
	"text/tabwriter"
	"time"
)

func WriteToConsole(stats []models.Stats) {
	tm.Clear()

	table := tm.NewTable(0, 10, 5, ' ', 0)
	fmt.Fprintf(table, "Container Name\tImage\tCPU %s\tMemory %s\tCurrent time:%s\n", "%", "%", time.Now().Format("15:04:05"))

	for _, stat := range stats {
		oneRow := fmt.Sprintf("%s\t%s\t%.2f\t%.2f\n", stat.Name, stat.Image, stat.CalculateCPUUsage(), stat.CalculateMemoryUsage())

		fmt.Fprint(table, oneRow)
	}

	tm.Print(table)

	tm.Flush()

}

func WriteTabWriter(stats []models.Stats) {

	table := tabwriter.NewWriter(os.Stdout, 0, 11, 5, ' ', 0)
	fmt.Fprintf(table, "\nContainer Name\tImage\tCPU %s\tMemory %s\n", "%", "%")

	for _, stat := range stats {
		oneRow := fmt.Sprintf("%s\t%s\t%.2f\t%.2f\n", stat.Name, stat.Image, stat.CalculateCPUUsage(), stat.CalculateMemoryUsage())
		fmt.Fprint(table, oneRow)
	}
	fmt.Fprint(table, "\n")
	table.Flush()
}
