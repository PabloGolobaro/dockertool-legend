package console

import (
	"fmt"
	tm "github.com/buger/goterm"
	pb "github.com/pablogolobaro/dockertool-legend/internal/api/protoStats"
)

func WriteToConsoleClient(stats *pb.Stats) {
	tm.Clear()

	table := tm.NewTable(0, 10, 5, ' ', 0)
	fmt.Fprintf(table, "\nContainer Name\tImage\tCPU %s\tMemory %s\n", "%", "%")

	for _, stat := range stats.ContainerStats {
		oneRow := fmt.Sprintf("%s\t%s\t%.2f\t%.2f\n", stat.Container.Name, stat.Container.Image, stat.CPU, stat.Memory)
		fmt.Fprint(table, oneRow)
	}

	tm.Print(table)

	tm.Flush()

}
