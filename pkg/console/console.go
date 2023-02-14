package console

import (
	"fmt"
	tm "github.com/buger/goterm"
	"os"
	"text/tabwriter"
	"time"
)

func WriteToConsole(stats []string) {
	tm.Clear()

	table := tm.NewTable(0, 11, 5, ' ', 0)
	fmt.Fprintf(table, "Container Name\tImage\tCPU %s\tMemory %s\tCurrent time:%s\n", "%", "%", time.Now().Format("15:04:05"))

	for _, stat := range stats {
		fmt.Fprint(table, stat)
	}

	tm.Print(table)

	tm.Flush()

}

func WriteOnce(stats []string) {

	table := tabwriter.NewWriter(os.Stdout, 0, 11, 5, ' ', 0)
	fmt.Fprintf(table, "\nContainer Name\tImage\tCPU %s\tMemory %s\n", "%", "%")

	for _, stat := range stats {
		fmt.Fprint(table, stat)
	}
	fmt.Fprint(table, "\n")
	table.Flush()
}
