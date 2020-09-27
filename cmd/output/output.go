package output

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"os"
)

func PrintResults(word string, out []string) {
	if len(out) < 1 {
		out = append(out, "Sorry, Goofi did not find anything!")
	}
	rows := make([]table.Row, len(out))
	for i := 0; i < len(out); i++ {
		rows[i] = table.Row{out[i]}
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{fmt.Sprintf("Result(s) for %s:", word)})
	t.AppendRows(rows)
	t.SetStyle(table.StyleLight)
	t.Render()
}
