package main

import (
	"fmt"
	tm "github.com/buger/goterm"
	"time"
)

func main() {
	started := 100
	finished := 250

	//tm.Clear() // Clear current screen
	for {
		tm.MoveCursorUp(7)
		// Clear current screen
		tm.Clear()
		totals := tm.NewTable(0, 10, 5, ' ', 0)

		// Based on http://golang.org/pkg/text/tabwriter

		fmt.Fprintf(totals, "Time\tStarted\tActive\tFinished\n")
		fmt.Fprintf(totals, "%s\t%d\t%d\t%d\n", "All", started, started-finished, finished)
		tm.Println(totals)

		tm.Flush()

		started--
		finished--
		time.Sleep(time.Second)

	}

}
