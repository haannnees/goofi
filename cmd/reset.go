package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goofi/cmd/csv"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:     "reset",
	Short:   "Remove all added abbreviations and reset goofi",
	Aliases: []string{"r"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return removeGoofiFile(goofiFilePath)
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}

func removeGoofiFile(path string) error {
	err := csv.RemoveCSVFile(path)
	if err != nil {
		return err
	}

	fmt.Print("Successfully deleted goofi file")

	return nil
}
