package cmd

import (
	"goofi/cmd/cli"
	csv2 "goofi/cmd/csv"
	"goofi/cmd/output"
	"strings"

	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:     "find",
	Short:   "Look up abbreviation",
	Aliases: []string{"f", "finder", "wtf"},
	RunE: func(cmd *cobra.Command, args []string) error {
		arg := cli.FlatArgs(args)
		data, err := csv2.OpenAndReadCSV(goofiFilePath)
		if err != nil {
			return err
		}
		return findResult(data, arg)
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
}

func findResult(data [][]string, arg string) error {
	var res []string
	for _, line := range data {
		if strings.ToLower(line[0]) == arg {
			res = append(res, line[1])
		}
	}

	output.PrintResults(arg, res)

	return nil
}

